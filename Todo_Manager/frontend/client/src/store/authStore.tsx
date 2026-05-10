import { createContext, useContext, useState, useEffect, type ReactNode } from "react"
import { login as apiLogin, apiLogout, register as apiRegister, getUserIdFromToken } from "@/api/Auth"

interface AuthState {
  token: string | null
  userId: number | null
  username: string | null
  isAuthenticated: boolean
  isLoading: boolean
  error: string | null
  login: (username: string, password: string) => Promise<void>
  register: (username: string, password: string) => Promise<void>
  logout: () => void
  clearError: () => void
}

const AuthContext = createContext<AuthState | null>(null)

export function AuthProvider({ children }: { children: ReactNode }) {
  const [token, setToken] = useState<string | null>(localStorage.getItem("token"))
  const [userId, setUserId] = useState<number | null>(() => {
    const t = localStorage.getItem("token")
    return t ? getUserIdFromToken(t) : null
  })
  const [username, setUsername] = useState<string | null>(localStorage.getItem("username"))
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    if (token) {
      localStorage.setItem("token", token)
    } else {
      localStorage.removeItem("token")
      localStorage.removeItem("username")
    }
  }, [token])

  const login = async (username: string, password: string) => {
    setIsLoading(true)
    setError(null)
    try {
      const data = await apiLogin(username, password)
      setToken(data.token)
      setUserId(getUserIdFromToken(data.token))
      setUsername(username)  // ← username это параметр функции login(username, password)
      localStorage.setItem("username", username)

      setUserId(getUserIdFromToken(data.token))
      setUsername(username)
      localStorage.setItem("username", username)
    } catch (err: unknown) {
      const e = err as { error: string }
      setError(e.error ?? "Ошибка входа")
      throw err
    } finally {
      setIsLoading(false)
    }
  }

  const register = async (username: string, password: string) => {
    setIsLoading(true)
    setError(null)
    try {
      await apiRegister(username, password)
      const data = await apiLogin(username, password)
      setToken(data.token)
      setUserId(getUserIdFromToken(data.token))
      setUsername(username)
      localStorage.setItem("username", username)
    } catch (err: unknown) {
      const e = err as { error: string }
      setError(e.error ?? "Ошибка регистрации")
      throw err
    } finally {
      setIsLoading(false)
    }
  }

  const logout = async () => {
    if (token) await apiLogout(token)
    setToken(null)
    setUserId(null)
    setUsername(null)
  }

  const clearError = () => setError(null)

  return (
    <AuthContext.Provider
      value={{
        token,
        userId,
        username,
        isAuthenticated: !!token,
        isLoading,
        error,
        login,
        register,
        logout,
        clearError,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}

// eslint-disable-next-line react-refresh/only-export-components
export function useAuth() {
  const ctx = useContext(AuthContext)
  if (!ctx) throw new Error("useAuth must be used within AuthProvider")
  return ctx
}
