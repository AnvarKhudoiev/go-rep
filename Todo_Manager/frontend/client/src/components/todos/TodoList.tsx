import { useState } from "react"
import { useTodos } from "@/store/todoStore"
import { Button } from "@/shadcn/components/ui/button"
import { EditTodoModal } from "@/components/todos/EditTodoModal"
import type { Todo } from "@/api/Todo"

const priorityColors: Record<string, string> = {
  low:    "bg-blue-100 text-blue-700",
  medium: "bg-yellow-100 text-yellow-700",
  high:   "bg-red-100 text-red-700",
}

const priorityLabels: Record<string, string> = {
  low: "Низкий", medium: "Средний", high: "Высокий",
}

export function TodoList({ onAddClick }: { onAddClick: () => void }) {
  const { todos, isLoading, toggleTodo, removeTodo } = useTodos()
  const [editingTodo, setEditingTodo] = useState<Todo | null>(null)

  const completed = todos.filter(t => t.completed)
  const pending = todos.filter(t => !t.completed)

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-40 text-muted-foreground text-sm">
        Загрузка задач...
      </div>
    )
  }

  return (
    <>
      <div className="space-y-4 max-w-2xl">
        <div className="flex items-center justify-between">
          <div>
            <h2 className="text-lg font-semibold">Задачи</h2>
            <p className="text-sm text-muted-foreground">
              {pending.length} активных · {completed.length} выполнено
            </p>
          </div>
          <Button size="sm" onClick={onAddClick}>+ Добавить задачу</Button>
        </div>

        {todos.length === 0 && (
          <div className="flex flex-col items-center justify-center h-40 rounded-xl border border-dashed border-border text-muted-foreground gap-2">
            <span className="text-2xl">✓</span>
            <p className="text-sm">Задач нет. Добавьте первую!</p>
          </div>
        )}

        {pending.length > 0 && (
          <div className="space-y-2">
            {pending.map(todo => (
              <TodoItem
                key={todo.id}
                todo={todo}
                onToggle={() => toggleTodo(todo.id, true)}
                onDelete={() => removeTodo(todo.id)}
                onEdit={() => setEditingTodo(todo)}
              />
            ))}
          </div>
        )}

        {completed.length > 0 && (
          <div className="space-y-2">
            <p className="text-xs font-medium text-muted-foreground uppercase tracking-wide">Выполнено</p>
            {completed.map(todo => (
              <TodoItem
                key={todo.id}
                todo={todo}
                onToggle={() => toggleTodo(todo.id, false)}
                onDelete={() => removeTodo(todo.id)}
                onEdit={() => setEditingTodo(todo)}
              />
            ))}
          </div>
        )}
      </div>

      {editingTodo && (
        <EditTodoModal
          todo={editingTodo}
          onClose={() => setEditingTodo(null)}
        />
      )}
    </>
  )
}

function TodoItem({
  todo,
  onToggle,
  onDelete,
  onEdit,
}: {
  todo: Todo
  onToggle: () => void
  onDelete: () => void
  onEdit: () => void
}) {
  const isOverdue = !todo.completed && todo.due_date && new Date(todo.due_date) < new Date()

  return (
    <div className={`flex items-start gap-3 p-3 rounded-lg border bg-card transition-opacity ${
      todo.completed ? "opacity-60" : ""
    } ${isOverdue ? "border-destructive/40" : "border-border"}`}>

      {/* Чекбокс */}
      <button
        onClick={onToggle}
        className={`mt-0.5 w-5 h-5 rounded-full border-2 shrink-0 flex items-center justify-center transition-colors ${
          todo.completed
            ? "bg-primary border-primary text-primary-foreground"
            : "border-muted-foreground hover:border-primary"
        }`}
      >
        {todo.completed && <span className="text-xs">✓</span>}
      </button>

      {/* Контент */}
      <div className="flex-1 min-w-0">
        <p className={`text-sm font-medium ${todo.completed ? "line-through text-muted-foreground" : ""}`}>
          {todo.title}
        </p>
        {todo.description && (
          <p className="text-xs text-muted-foreground mt-0.5 truncate">{todo.description}</p>
        )}
        <div className="flex items-center gap-2 mt-1.5 flex-wrap">
          {/* Приоритет */}
          {todo.priority && (
            <span className={`text-xs px-2 py-0.5 rounded-full font-medium ${priorityColors[todo.priority]}`}>
              {priorityLabels[todo.priority]}
            </span>
          )}
          {/* Дедлайн */}
          {todo.due_date && (
            <span className={`text-xs ${isOverdue ? "text-destructive font-medium" : "text-muted-foreground"}`}>
              {isOverdue ? "⚠ " : ""}
              {new Date(todo.due_date).toLocaleDateString("ru", {
                day: "numeric", month: "short",
              })}
            </span>
          )}
          {/* Категория */}
          {todo.category?.name && (
            <span
              className="text-xs px-2 py-0.5 rounded-full font-medium"
              style={{
                backgroundColor: todo.category.color + "22",
                color: todo.category.color,
              }}
            >
              {todo.category.name}
            </span>
          )}
          {/* Теги */}
          {todo.tags?.map(tag => (
            <span key={tag.id} className="text-xs px-2 py-0.5 rounded-full bg-muted text-muted-foreground">
              #{tag.name}
            </span>
          ))}
        </div>
      </div>

      {/* Кнопки */}
      <div className="flex items-center gap-1 shrink-0">
        <button
          onClick={onEdit}
          className="text-muted-foreground hover:text-foreground transition-colors p-1 rounded"
          title="Редактировать"
        >
          ✎
        </button>
        <button
          onClick={onDelete}
          className="text-muted-foreground hover:text-destructive transition-colors p-1 rounded"
          title="Удалить"
        >
          ×
        </button>
      </div>
    </div>
  )
}