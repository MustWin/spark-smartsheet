import React from "react"
import ReactDOM from "react-dom"
import {createStore} from "redux"
import MuiThemeProvider from "material-ui/styles/MuiThemeProvider";
import Paper from "material-ui/Paper"

import injectTapEventPlugin from "react-tap-event-plugin";

import Head from "./Head.js"
import Login from "./Login.js"
import Reducer from "./Reducer.js"

const store = createStore(Reducer);

const MOUNT_NODE = document.getElementById("root");

var render = function () {
  const style = {
    "padding": "2em",
    "minHeight": "95%"
  };

  ReactDOM.render(
    <MuiThemeProvider>
      <Paper style={style} zDepth={3}>
        <Head />
        <Login store={store} />
      </Paper>
    </MuiThemeProvider>,
    MOUNT_NODE);
};

injectTapEventPlugin();
render();
store.subscribe(render);
