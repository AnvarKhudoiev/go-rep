import { useState } from "react"
import { useTodos } from "@/store/todoStore"
import { Button } from "@/shadcn/components/ui/button"
import { Input } from "@/shadcn/components/ui/input"

export function AddTodoPage({ onSuccess }: { onSuccess: () => void }) {
  const { addTodo, categories } = useTodos()
  const [title, setTitle] = useState("")
  const [description, setDescription] = useState("")
  const [categoryId, setCategoryId] = useState<number | undefined>()
  const [tagInput, setTagInput] = useState("")
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const selectedCategory = categories.find(c => c.id === categoryId)

  const handleSubmit = async () => {
    if (!title.trim()) return
    setIsLoading(true)
    setError(null)
    try {
      const tags = tagInput
        .split(",")
        .map(t => t.trim())
        .filter(Boolean)
        .map(name => ({ name }))

      await addTodo({
        title: title.trim(),
        description: description.trim(),
        ...(categoryId ? { category_id: categoryId } : {}),
        tags: tags.length > 0 ? tags : undefined,
      })

      onSuccess()
    } catch {
      setError("Ошибка при создании задачи")
    } finally {
      setIsLoading(false)
    }
  }

  const defaultCategories = categories.filter(c => c.is_default)
  const userCategories = categories.filter(c => !c.is_default)

  return (
    <div className="max-w-lg space-y-6">
      <div>
        <h2 className="text-xl font-semibold">Новая задача</h2>
        <p className="text-sm text-muted-foreground mt-1">Заполните детали задачи</p>
      </div>

      <div className="space-y-4">
        <div className="space-y-1.5">
          <label className="text-sm font-medium">Название <span className="text-destructive">*</span></label>
          <Input
            placeholder="Что нужно сделать?"
            value={title}
            onChange={e => setTitle(e.target.value)}
            onKeyDown={e => e.key === "Enter" && handleSubmit()}
            autoFocus
          />
        </div>

        <div className="space-y-1.5">
          <label className="text-sm font-medium">Описание</label>
          <textarea
            placeholder="Подробное описание задачи..."
            value={description}
            onChange={e => setDescription(e.target.value)}
            rows={4}
            className="w-full rounded-md border border-input bg-background px-3 py-2 text-sm placeholder:text-muted-foreground focus:outline-none focus:ring-1 focus:ring-ring resize-none"
          />
        </div>

        <div className="space-y-1.5">
          <label className="text-sm font-medium">Категория</label>
          <div className="relative">
            <select
              value={categoryId ?? ""}
              onChange={e => setCategoryId(e.target.value ? Number(e.target.value) : undefined)}
              className="w-full h-9 rounded-md border border-input bg-background pl-8 pr-3 text-sm focus:outline-none focus:ring-1 focus:ring-ring appearance-none"
            >
              <option value="">Без категории</option>
              {defaultCategories.length > 0 && (
                <optgroup label="Общие">
                  {defaultCategories.map(cat => (
                    <option key={cat.id} value={cat.id}>{cat.name}</option>
                  ))}
                </optgroup>
              )}
              {userCategories.length > 0 && (
                <optgroup label="Мои">
                  {userCategories.map(cat => (
                    <option key={cat.id} value={cat.id}>{cat.name}</option>
                  ))}
                </optgroup>
              )}
            </select>
            {/* Цветной индикатор выбранной категории */}
            <div
              className="absolute left-2.5 top-1/2 -translate-y-1/2 w-3 h-3 rounded-full transition-colors"
              style={{ backgroundColor: selectedCategory?.color ?? "#e5e7eb" }}
            />
          </div>
        </div>

        <div className="space-y-1.5">
          <label className="text-sm font-medium">Теги</label>
          <Input
            placeholder="работа, важное, дом (через запятую)"
            value={tagInput}
            onChange={e => setTagInput(e.target.value)}
          />
          <p className="text-xs text-muted-foreground">Введите теги через запятую</p>
        </div>

        {error && (
          <div className="text-sm text-destructive bg-destructive/10 border border-destructive/20 px-3 py-2 rounded-lg">
            {error}
          </div>
        )}

        <div className="flex gap-3 pt-2">
          <Button
            onClick={handleSubmit}
            disabled={isLoading || !title.trim()}
            className="flex-1"
          >
            {isLoading ? "Создание..." : "Создать задачу"}
          </Button>
          <Button
            variant="outline"
            onClick={onSuccess}
            disabled={isLoading}
          >
            Отмена
          </Button>
        </div>
      </div>
    </div>
  )
}