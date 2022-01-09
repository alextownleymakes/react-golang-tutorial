import './TodoItem.scss'

export const TodoItem = (o) => {
    return(
        <article className="todo-item">
            <section>
                <h3>{o.title}</h3>
                <p>{o.description}</p>
            </section>
            <input type="checkbox" defaultChecked={o.isCompleted} />
        </article>
    )
}