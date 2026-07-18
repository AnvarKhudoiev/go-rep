const TODO_URL = "http://localhost:8082"

const authHeaders = () => ({
  "Content-Type": "application/json",
  Authorization: `Bearer ${localStorage.getItem("token") ?? ""}`,
})

export interface Tag {
  id: number
  name: string
}

export interface Category {
  id: number
  name: string
  color: string
  user_id: number
  is_default: boolean  // ← добавить
}

export interface Todo {
  id: number
  title: string
  description: string
  completed: boolean
  user_id: number
  category_id: number
  category: Category
  tags: Tag[]
  created_at: string
  updated_at: string
  priority: "low" | "medium" | "high"  // ← добавить
  due_date: string | null               // ← добавить
}

export interface CreateTodoInput {
  title: string
  description?: string
  completed?: boolean
  category_id?: number
  priority?: "low" | "medium" | "high"  // ← добавить
  due_date?: string                      // ← добавить
  tags?: { name: string }[]
}

export const getTodos = async (): Promise<Todo[]> => {
  const res = await fetch(`${TODO_URL}/api/todos`, { headers: authHeaders() })
  if (!res.ok) throw await res.json()
  return res.json()
}

export const createTodo = async (input: CreateTodoInput): Promise<Todo> => {
  const res = await fetch(`${TODO_URL}/api/todos`, {
    method: "POST",
    headers: authHeaders(),
    body: JSON.stringify(input),
  })
  if (!res.ok) throw await res.json()
  return res.json()
}

export const updateTodo = async (id: number, input: Partial<CreateTodoInput & { completed: boolean }>): Promise<Todo> => {
  const res = await fetch(`${TODO_URL}/api/todos/${id}`, {
    method: "PUT",
    headers: authHeaders(),
    body: JSON.stringify(input),
  })
  if (!res.ok) throw await res.json()
  return res.json()
}

export const deleteTodo = async (id: number): Promise<void> => {
  const res = await fetch(`${TODO_URL}/api/todos/${id}`, {
    method: "DELETE",
    headers: authHeaders(),
  })
  if (!res.ok) throw await res.json()
}

export const getCategories = async (): Promise<Category[]> => {
  const res = await fetch(`${TODO_URL}/api/categories`, { headers: authHeaders() })
  if (!res.ok) throw await res.json()
  return res.json()
}

export interface DayStats {
  date: string
  label: string
  created: number
  completed: number
}

export interface CategoryStats {
  name: string
  color: string
  count: number
}

export interface Stats {
  total: number
  completed: number
  active: number
  overdue: number        // ← добавить
  completion_rate: number
  week: DayStats[]
  by_category: CategoryStats[]
}

export const getStats = async (): Promise<Stats> => {
  const res = await fetch(`${TODO_URL}/api/stats`, { headers: authHeaders() })
  if (!res.ok) throw await res.json()
  return res.json()
}