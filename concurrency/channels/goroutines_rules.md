# Goroutines Rules

1. O canal deve ser fechado pelo remetente, nunca pelo receptor.
2. O canal deve ser fechado quando o remetente não tem mais dados para enviar: `defer close(ch)` dentro da goroutine.
3. O canal nunca deve ser fechado mais de uma vez.
4. Um for loop com select deve ter uma condição de parada para evitar deadlock `for ch != nil {}`.