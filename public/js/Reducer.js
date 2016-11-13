import _ from "underscore";

function Reducer(state = {}, action) {
  console.log("REDUCER", action.type, action);
  var o = Object.assign({}, state);
  switch (action.type) {
    case "LOGIN":
      o.token = action.result;
      if (action.result === "TOKEN") {o.loggedIn = true}
      console.log("SENDING",o);
      return o;
    case "LOGOUT":
      o.token = "";
      o.loggedIn = false;
      return o;
    default:
      return state;
  }
}

export default Reducer
