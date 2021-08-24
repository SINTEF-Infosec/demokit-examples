import './App.css';

const LIGHT_NODE_URL = "http://localhost:8081/"
const LIGHT_ON_ENDPOINT = "turn_light_on"
const LIGHT_OFF_ENDPOINT = "turn_light_off"

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <button onClick={() => fetch(LIGHT_NODE_URL + LIGHT_ON_ENDPOINT)} style={{margin: "2em"}}>Turn light on</button>
        <button onClick={() => fetch(LIGHT_NODE_URL + LIGHT_OFF_ENDPOINT)} style={{margin: "2em"}}>Turn light off</button>
      </header>
    </div>
  );
}

export default App;
