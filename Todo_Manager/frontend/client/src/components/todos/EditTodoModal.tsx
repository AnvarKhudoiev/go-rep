import { useState, useEffect } from "react"
import { useTodos } from "@/store/todoStore"
import { Button } from "@/shadcn/components/ui/button"
import { Input } from "@/shadcn/components/ui/input"
import type { Todo } from "@/api/Todo"

type Priority = "low" | "medium" | "high"

const priorityConfig: Record<Priority, { label: string; color: string }> = {
  low:    { label: "Низкий",  color: "bg-blue-100 text-blue-700 border-blue-200" },
  medium: { label: "Средний", color: "bg-yellow-100 text-yellow-700 border-yellow-200" },
  high:   { label: "Высокий", color: "bg-red-100 text-red-700 border-red-200" },
}

interface EditTodoModalProps {
  todo: Todo
  onClose: () => void
}

export function EditTodoModal({ todo, onClose }: EditTodoModalProps) {
  const { editTodo, categories } = useTodos()
  const [title, setTitle] = useState(todo.title)
  const [description, setDescription] = useState(todo.description)
  const [categoryId, setCategoryId] = useState<number | undefined>(
    todo.category_id ?? undefined
  )
  const [priority, setPriority] = useState<Priority>((todo.priority as Priority) ?? "medium")
  const [dueDate, setDueDate] = useState(
    todo.due_date ? new Date(todo.due_date).toISOString().slice(0, 16) : ""
  )
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const selectedCategory = categories.find(c => c.id === categoryId)
  const defaultCategories = categories.filter(c => c.is_default)
  const userCategories = categories.filter(c => !c.is_default)

  // Закрыть по Escape
  useEffect(() => {
    const handler = (e: KeyboardEvent) => e.key === "Escape" && onClose()
    window.addEventListener("keydown", handler)
    return () => window.removeEventListener("keydown", handler)
  }, [onClose])

  const handleSubmit = async () => {
    if (!title.trim()) return
    setIsLoading(true)
    setError(null)
    try {
      await editTodo(todo.id, {
        title: title.trim(),
        description: description.trim(),
        ...(categoryId ? { category_id: categoryId } : {}),
        priority,
        ...(dueDate ? { due_date: new Date(dueDate).toISOString() } : {}),
      })
      onClose()
    } catch {
      setError("Ошибка при сохранении")
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* Overlay */}
      <div
        className="absolute inset-0 bg-black/40 backdrop-blur-sm"
        onClick={onClose}
      />

      {/* Modal */}
      <div className="relative z-10 w-full max-w-md bg-card border border-border rounded-xl shadow-lg p-6 space-y-5 mx-4">
        <div className="flex items-center justify-between">
          <h2 className="text-lg font-semibold">Редактировать задачу</h2>
          <button
            onClick={onClose}
            className="text-muted-foreground hover:text-foreground transition-colors text-xl leading-none"
          >
            ×
          </button>
        </div>

        <div className="space-y-4">
          {/* Название */}
          <div className="space-y-1.5">
            <label className="text-sm font-medium">Название</label>
            <Input
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
                className="absolute left-2.5 top-1/2 -translate-y-1/2 w-3 h-3 rounded-full"
                style={{ backgroundColor: selectedCategory?.color ?? "#e5e7eb" }}
              />
            </div>
          </div>

          {error && (
            <p className="text-sm text-destructive bg-destructive/10 border border-destructive/20 px-3 py-2 rounded-lg">
              {error}
            </p>
          )}

          <div className="flex gap-3 pt-1">
            <Button onClick={handleSubmit} disabled={isLoading || !title.trim()} className="flex-1">
              {isLoading ? "Сохранение..." : "Сохранить"}
            </Button>
            <Button variant="outline" onClick={onClose} disabled={isLoading}>
              Отмена
            </Button>
          </div>
        </div>
      </div>
    </div>
  )
}