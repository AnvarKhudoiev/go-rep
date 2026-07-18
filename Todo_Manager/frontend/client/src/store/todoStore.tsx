import { createContext, useContext, useState, useEffect, type ReactNode } from "react"
import {
  getTodos,
  createTodo,
  updateTodo,
  deleteTodo,
  getCategories,
  type Todo,
  type Category,
  type CreateTodoInput,
} from "@/api/Todo"

interface TodoState {
  todos: Todo[]
  categories: Category[]
  isLoading: boolean
  error: string | null
  fetchTodos: () => Promise<void>
  addTodo: (input: CreateTodoInput) => Promise<void>
  toggleTodo: (id: number, completed: boolean) => Promise<void>
  removeTodo: (id: number) => Promise<void>
  fetchCategories: () => Promise<void>
  editTodo: (id: number, input: Partial<CreateTodoInput>) => Promise<void>
}

const TodoContext = createContext<TodoState | null>(null)

export function TodoProvider({ children }: { children: ReactNode }) {
  const [todos, setTodos] = useState<Todo[]>([])
  const [categories, setCategories] = useState<Category[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const fetchTodos = async () => {
    setIsLoading(true)
    try {
      const data = await getTodos()
      setTodos(data)
    } catch {
      setError("Ошибка загрузки задач")
    } finally {
      setIsLoading(false)
    }
  }

  const fetchCategories = async () => {
    try {
      const data = await getCategories()
      setCategories(data)
    } catch {
      setError("Ошибка загрузки категорий")
    }
  }

  const addTodo = async (input: CreateTodoInput) => {
    const todo = await createTodo(input)
    setTodos(prev => [todo, ...prev])
  }

  const toggleTodo = async (id: number, completed: boolean) => {
    const updated = await updateTodo(id, { completed })
    setTodos(prev => prev.map(t => t.id === id ? updated : t))
  }

  const removeTodo = async (id: number) => {
    await deleteTodo(id)
    setTodos(prev => prev.filter(t => t.id !== id))
  }

  const editTodo = async (id: number, input: Partial<CreateTodoInput>) => {
    const updated = await updateTodo(id, input)
    setTodos(prev => prev.map(t => t.id === id ? updated : t))
  }

  useEffect(() => {
    // eslint-disable-next-line react-hooks/set-state-in-effect
    fetchTodos()
    fetchCategories()
  }, [])

  return (
    <TodoContext.Provider value={{
      todos,
      categories,
      isLoading,
      error,
      fetchTodos,
      addTodo,
      toggleTodo,
      removeTodo,
      fetchCategories,
      editTodo,
    }}>
      {children}
    </TodoContext.Provider>
  )
}

// eslint-disable-next-line react-refresh/only-export-components
export function useTodos() {
  const ctx = useContext(TodoContext)
  if (!ctx) throw new Error("useTodos must be used within TodoProvider")
  return ctx
}
