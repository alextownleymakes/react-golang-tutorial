import { TodoItem } from "./TodoItem"


export const ToDoList = ({ todos }) => {
    return (
        <ul className="todo-list">
            {
                todos.map((todo, i) => {
                    return <li key={i}><TodoItem {...todo}/></li>
                })
            }
        </ul>
    )
}