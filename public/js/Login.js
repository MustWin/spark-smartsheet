import React from "react"
import RaisedButton from "material-ui/RaisedButton"
import TextField from "material-ui/TextField";
import {connect} from "react-redux";

class Login extends React.Component {
  constructor(props) {
    super(props);
    const {store} = props;
    this.dispatcher = store.dispatch;
    this.onLoginClick = this.onLoginClick.bind(this);
    this.onLogoutClick = this.onLogoutClick.bind(this);
    this.state = Object.assign({token: "", loggedIn: false}, store.getState());
  }

  onLoginClick() { this.dispatcher({type: "LOGIN", result: this.state.token}) }

  onLogoutClick() { this.dispatcher({type: "LOGOUT"}) }

  componentWillReceiveProps(newProps) { this.setState({token: newProps.token, loggedIn: newProps.loggedIn}); }

  render() {
    if (this.state.loggedIn) {
      return (
        <div>
          <h4>Logged In</h4>
          <RaisedButton label={"Logout"} onClick={this.onLogoutClick} />
        </div>
      )
    }
    return (
      <div>
        <TextField hintText={"login token"} value={this.state.token}
                   onChange={(i)=> this.setState({token: i.target.value})} /><br />
        <RaisedButton label={"Login"} onClick={this.onLoginClick} />
      </div>
    );
  }
}

const mapStateToProps = (state) => { return Object.assign({}, state, {}); };
export default connect(mapStateToProps)(Login);
