import React from "react"

class Hello extends React.Component {

  constructor(value) {
    super(value);
    this.props = value;
  }

  render() {
    const {value, onTimer, onReset} = this.props;
    return (
      <div>
        <h3>Hi! {value.name}</h3>
        <button onClick={onTimer}>Go</button>
        <button onClick={onReset}>Reset</button>
      </div>
    );
  }
}

export default Hello
