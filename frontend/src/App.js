
import './App.css';
import { Header } from './components/Header';
import { ToDoList } from './components/ToDoList.js';

const todos = [
  {title: "Laundry", description: 'take some to your mom\'s', isCompleted: false},
  {title: "Dishes", description: 'just wash them', isCompleted: false},
  {title: "Organize music room", description: 'gotta get rid of that couch', isCompleted: false},
  {title: "Empty boxes", description: 'one at a time', isCompleted: false},
  {title: "Hi five sam", description: 'he deserves it', isCompleted: false}
]
function App() {
  return (
    <div className="App">
      <Header />
      <ToDoList todos={todos} />
    </div>
  );
}

export default App;
