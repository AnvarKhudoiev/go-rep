import { useState } from "react"
import { useTodos } from "@/store/todoStore"
import { Button } from "@/shadcn/components/ui/button"
import { Input } from "@/shadcn/components/ui/input"

type Priority = "low" | "medium" | "high"

const priorityConfig: Record<Priority, { label: string; color: string }> = {
  low:    { label: "Низкий",   color: "bg-blue-100 text-blue-700 border-blue-200" },
  medium: { label: "Средний",  color: "bg-yellow-100 text-yellow-700 border-yellow-200" },
  high:   { label: "Высокий",  color: "bg-red-100 text-red-700 border-red-200" },
}

export function AddTodoPage({ onSuccess }: { onSuccess: () => void }) {
  const { addTodo, categories } = useTodos()
  const [title, setTitle] = useState("")
  const [description, setDescription] = useState("")
  const [categoryId, setCategoryId] = useState<number | undefined>()
  const [tagInput, setTagInput] = useState("")
  const [priority, setPriority] = useState<Priority>("medium")
  const [dueDate, setDueDate] = useState("")
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const selectedCategory = categories.find(c => c.id === categoryId)
  const defaultCategories = categories.filter(c => c.is_default)
  const userCategories = categories.filter(c => !c.is_default)

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
        priority,
        ...(dueDate ? { due_date: new Date(dueDate).toISOString() } : {}),
        tags: tags.length > 0 ? tags : undefined,
      })

      onSuccess()
    } catch {
      setError("Ошибка при создании задачи")
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="max-w-lg space-y-6">
      <div>
        <h2 className="text-xl font-semibold">Новая задача</h2>
        <p className="text-sm text-muted-foreground mt-1">Заполните детали задачи</p>
      </div>

      <div className="space-y-4">

        {/* Название */}
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

        {/* Описание */}
        <div className="space-y-1.5">
          <label className="text-sm font-medium">Описание</label>
          <textarea
            placeholder="Подробное описание задачи..."
            value={description}
            onChange={e => setDescription(e.target.value)}
            rows={3}
            className="w-full rounded-md border border-input bg-background px-3 py-2 text-sm placeholder:text-muted-foreground focus:outline-none focus:ring-1 focus:ring-ring resize-none"
          />
        </div>

        {/* Приоритет */}
        <div className="space-y-1.5">
          <label className="text-sm font-medium">Приоритет</label>
          <div className="flex gap-2">
            {(["low", "medium", "high"] as Priority[]).map(p => (
              <button
                key={p}
                onClick={() => setPriority(p)}
                className={`flex-1 py-1.5 rounded-md text-sm font-medium border transition-all ${
                  priority === p
                    ? priorityConfig[p].color
                    : "border-border text-muted-foreground hover:border-foreground"
                }`}
              >
                {priorityConfig[p].label}
              </button>
            ))}
          </div>
        </div>

        {/* Дедлайн */}
        <div className="space-y-1.5">
          <label className="text-sm font-medium">Дедлайн</label>
          <input
            type="datetime-local"
            value={dueDate}
            onChange={e => setDueDate(e.target.value)}
            className="w-full h-9 rounded-md border border-input bg-background px-3 text-sm focus:outline-none focus:ring-1 focus:ring-ring"
          />
          {dueDate && (
            <p className="text-xs text-muted-foreground">
              {new Date(dueDate).toLocaleString("ru", {
                day: "numeric", month: "long", year: "numeric",
                hour: "2-digit", minute: "2-digit"
              })}
            </p>
          )}
        </div>

        {/* Категория */}
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
            <div
              className="absolute left-2.5 top-1/2 -translate-y-1/2 w-3 h-3 rounded-full transition-colors"
              style={{ backgroundColor: selectedCategory?.color ?? "#e5e7eb" }}
            />
          </div>
        </div>

        {/* Теги */}
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
          <Button onClick={handleSubmit} disabled={isLoading || !title.trim()} className="flex-1">
            {isLoading ? "Создание..." : "Создать задачу"}
          </Button>
          <Button variant="outline" onClick={onSuccess} disabled={isLoading}>
            Отмена
          </Button>
        </div>
      </div>
    </div>
  )
}