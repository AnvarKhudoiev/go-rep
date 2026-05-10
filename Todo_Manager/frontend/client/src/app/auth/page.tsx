import { useState } from "react"
import { useAuth } from "@/store/authStore"
import { Button } from "@/shadcn/components/ui/button"
import { Input } from "@/shadcn/components/ui/input"

export default function AuthPage() {
  const { login, register, isLoading, error, clearError } = useAuth()
  const [mode, setMode] = useState<"login" | "register">("login")
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  const handleSubmit = async () => {
    if (!username || !password) return
    try {
      if (mode === "login") {
        await login(username, password)
      } else {
        await register(username, password)
      }
    } catch {
      // ошибка уже в store
    }
  }

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") handleSubmit()
  }

  const switchMode = (newMode: "login" | "register") => {
    if (newMode === mode) return
    clearError()
    setUsername("")
    setPassword("")
    setMode(newMode)
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-muted/40">
      <div className="w-full max-w-sm px-4">

        <div className="bg-card border border-border rounded-xl p-8 space-y-6">

          {/* Шапка */}
          <div className="space-y-1">
            <div className="flex items-center gap-2.5 mb-5">
              <div className="w-8 h-8 rounded-lg bg-foreground flex items-center justify-center text-background text-sm font-bold">
                ✓
              </div>
              <span className="text-sm font-medium text-muted-foreground">Todo Manager</span>
            </div>
            <h1 className="text-2xl font-semibold tracking-tight">
              {mode === "login" ? "Добро пожаловать" : "Создать аккаунт"}
            </h1>
            <p className="text-sm text-muted-foreground">
              {mode === "login"
                ? "Войдите в свой аккаунт чтобы продолжить"
                : "Зарегистрируйтесь чтобы начать"}
            </p>
          </div>

          {/* Переключатель */}
          <div className="flex bg-muted rounded-lg p-1 gap-1">
            <button
              onClick={() => switchMode("login")}
              className={`flex-1 py-1.5 text-sm font-medium rounded-md transition-all ${
                mode === "login"
                  ? "bg-background text-foreground border border-border"
                  : "text-muted-foreground hover:text-foreground"
              }`}
            >
              Войти
            </button>
            <button
              onClick={() => switchMode("register")}
              className={`flex-1 py-1.5 text-sm font-medium rounded-md transition-all ${
                mode === "register"
                  ? "bg-background text-foreground border border-border"
                  : "text-muted-foreground hover:text-foreground"
              }`}
            >
              Регистрация
            </button>
          </div>

          {/* Поля */}
          <div className="space-y-3">
            <div className="space-y-1.5">
              <label className="text-sm font-medium">Имя пользователя</label>
              <Input
                placeholder="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                onKeyDown={handleKeyDown}
                disabled={isLoading}
                autoFocus
              />
            </div>

            <div className="space-y-1.5">
              <label className="text-sm font-medium">Пароль</label>
              <Input
                type="password"
                placeholder="••••••••"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                onKeyDown={handleKeyDown}
                disabled={isLoading}
              />
            </div>

            {error && (
              <div className="text-sm text-destructive bg-destructive/10 border border-destructive/20 px-3 py-2 rounded-lg">
                {error}
              </div>
            )}

            <Button
              className="w-full mt-1"
              onClick={handleSubmit}
              disabled={isLoading || !username || !password}
            >
              {isLoading
                ? "Загрузка..."
                : mode === "login"
                ? "Войти"
                : "Зарегистрироваться"}
            </Button>
          </div>
        </div>

        <p className="text-center text-sm text-muted-foreground mt-4">
          {mode === "login" ? "Нет аккаунта?" : "Уже есть аккаунт?"}{" "}
          <button
            onClick={() => switchMode(mode === "login" ? "register" : "login")}
            className="text-foreground font-medium underline-offset-4 hover:underline"
          >
            {mode === "login" ? "Зарегистрироваться" : "Войти"}
          </button>
        </p>
      </div>
    </div>
  )
}
