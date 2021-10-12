import React from "react";
import { 
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom"
import {HistoryPricesScreen} from "./screens/HistoryPricesScreen";
import { HomeScreen } from "./screens/HomeScreen";
import "./App.css";

function App() {  
  return (
    <Router>
      <div className="App">
        <HomeScreen />
        <Switch>
          <Route path="/historyPrices/:id">
            <HistoryPricesScreen />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
