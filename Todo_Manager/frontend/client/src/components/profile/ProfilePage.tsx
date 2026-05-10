import { useEffect, useState } from "react"
import { useAuth } from "@/store/authStore"
import { getStats, type Stats } from "@/api/Todo"
import {
  PieChart, Pie, Cell,
  BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer,
} from "recharts"

export function ProfilePage() {
  const { username } = useAuth()
  const [stats, setStats] = useState<Stats | null>(null)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    getStats()
      .then(setStats)
      .catch(() => {})
      .finally(() => setIsLoading(false))
  }, [])

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-40 text-muted-foreground text-sm">
        Загрузка статистики...
      </div>
    )
  }

  if (!stats) {
    return (
      <div className="flex items-center justify-center h-40 text-muted-foreground text-sm">
        Не удалось загрузить статистику
      </div>
    )
  }

  return (
    <div className="max-w-3xl space-y-8">

      {/* Шапка */}
      <div className="flex items-center gap-4">
        <div className="w-16 h-16 rounded-full bg-foreground flex items-center justify-center text-background text-2xl font-semibold">
          {username?.slice(0, 2).toUpperCase() ?? "AN"}
        </div>
        <div>
          <h2 className="text-2xl font-semibold">{username}</h2>
          <p className="text-sm text-muted-foreground">Личный профиль</p>
        </div>
      </div>

      {/* Счётчики */}
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        {[
          { label: "Всего задач",  value: stats.total,           color: "text-foreground"  },
          { label: "Активных",     value: stats.active,          color: "text-blue-500"    },
          { label: "Выполнено",    value: stats.completed,       color: "text-green-500"   },
          { label: "Прогресс",     value: `${stats.completion_rate}%`, color: "text-orange-500" },
        ].map(item => (
          <div key={item.label} className="rounded-xl border border-border bg-card p-4 space-y-1">
            <p className="text-xs text-muted-foreground">{item.label}</p>
            <p className={`text-3xl font-bold ${item.color}`}>{item.value}</p>
          </div>
        ))}
      </div>

      {/* Bar chart — 7 дней */}
      <div className="rounded-xl border border-border bg-card p-5 space-y-3">
        <h3 className="text-sm font-medium">Активность за 7 дней</h3>
        {stats.total === 0 ? (
          <p className="text-sm text-muted-foreground py-8 text-center">Нет данных</p>
        ) : (
          <ResponsiveContainer width="100%" height={200}>
            <BarChart data={stats.week} barGap={4}>
              <XAxis dataKey="label" tick={{ fontSize: 12 }} axisLine={false} tickLine={false} />
              <YAxis allowDecimals={false} tick={{ fontSize: 12 }} axisLine={false} tickLine={false} width={24} />
              <Tooltip
                contentStyle={{ fontSize: 12, borderRadius: 8, border: "1px solid var(--border)" }}
                labelStyle={{ fontWeight: 500 }}
              />
              <Bar dataKey="created"   name="Создано"    fill="#94a3b8" radius={[4,4,0,0]} />
              <Bar dataKey="completed" name="Выполнено"  fill="#22c55e" radius={[4,4,0,0]} />
            </BarChart>
          </ResponsiveContainer>
        )}
      </div>

      {/* Donut — по категориям */}
      <div className="rounded-xl border border-border bg-card p-5 space-y-3">
        <h3 className="text-sm font-medium">Задачи по категориям</h3>
        {stats.by_category.length === 0 ? (
          <p className="text-sm text-muted-foreground py-8 text-center">Нет данных</p>
        ) : (
          <div className="flex items-center gap-8">
            <ResponsiveContainer width={180} height={180}>
              <PieChart>
                <Pie
                  data={stats.by_category}
                  dataKey="count"
                  nameKey="name"
                  cx="50%" cy="50%"
                  innerRadius={50} outerRadius={80}
                  paddingAngle={3}
                >
                  {stats.by_category.map((entry, i) => (
                    <Cell key={i} fill={entry.color} />
                  ))}
                </Pie>
                <Tooltip contentStyle={{ fontSize: 12, borderRadius: 8 }} />
              </PieChart>
            </ResponsiveContainer>
            <div className="space-y-2 flex-1">
              {stats.by_category.map((cat, i) => (
                <div key={i} className="flex items-center justify-between text-sm">
                  <div className="flex items-center gap-2">
                    <div className="w-3 h-3 rounded-full shrink-0" style={{ backgroundColor: cat.color }} />
                    <span className="text-muted-foreground">{cat.name}</span>
                  </div>
                  <span className="font-medium">{cat.count}</span>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>

    </div>
  )
}