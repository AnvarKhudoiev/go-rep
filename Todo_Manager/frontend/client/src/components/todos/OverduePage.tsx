import { useTodos } from "@/store/todoStore"
import type { Todo } from "@/api/Todo"

export function OverduePage() {
  const { todos, toggleTodo, removeTodo } = useTodos()

  const now = new Date()
  const overdue = todos.filter(t =>
    !t.completed && t.due_date && new Date(t.due_date) < now
  )

  if (overdue.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center h-40 text-muted-foreground gap-2">
        <span className="text-2xl">✓</span>
        <p className="text-sm">Просроченных задач нет!</p>
      </div>
    )
  }

  return (
    <div className="max-w-2xl space-y-4">
      <div>
        <h2 className="text-lg font-semibold">Просроченные задачи</h2>
        <p className="text-sm text-muted-foreground">{overdue.length} задач требуют внимания</p>
      </div>

      <div className="space-y-2">
        {overdue.map(todo => (
          <OverdueItem
            key={todo.id}
            todo={todo}
            onComplete={() => toggleTodo(todo.id, true)}
            onDelete={() => removeTodo(todo.id)}
          />
        ))}
      </div>
    </div>
  )
}

function OverdueItem({
  todo,
  onComplete,
  onDelete,
}: {
  todo: Todo
  onComplete: () => void
  onDelete: () => void
}) {
  const daysOverdue = todo.due_date
    // eslint-disable-next-line react-hooks/purity
    ? Math.floor((Date.now() - new Date(todo.due_date).getTime()) / 86400000)
    : 0

  return (
    <div className="flex items-start gap-3 p-3 rounded-lg border border-destructive/30 bg-destructive/5">
      <div className="flex-1 min-w-0">
        <p className="text-sm font-medium">{todo.title}</p>
        {todo.description && (
          <p className="text-xs text-muted-foreground mt-0.5 truncate">{todo.description}</p>
        )}
        <div className="flex items-center gap-2 mt-1.5">
          <span className="text-xs text-destructive font-medium">
            Просрочено на {daysOverdue === 0 ? "сегодня" : `${daysOverdue} дн.`}
          </span>
          {todo.due_date && (
            <span className="text-xs text-muted-foreground">
              · {new Date(todo.due_date).toLocaleDateString("ru", {
                day: "numeric", month: "short", hour: "2-digit", minute: "2-digit"
              })}
            </span>
          )}
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
        </div>
      </div>

      <div className="flex items-center gap-2 shrink-0">
        <button
          onClick={onComplete}
          className="text-xs px-2 py-1 rounded-md border border-green-500 text-green-600 hover:bg-green-50 transition-colors"
        >
          Выполнить
        </button>
        <button
          onClick={onDelete}
          className="text-muted-foreground hover:text-destructive transition-colors text-lg leading-none"
        >
          ×
        </button>
      </div>
    </div>
  )
}