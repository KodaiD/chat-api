import React, { Component } from "react";
import "./Message.scss";

class Message extends Component {
  constructor(props) {
    super(props);
    let temp = JSON.parse(this.props.message);
    this.state = {
      message: temp
    };
  }

  render() {
    const userId = this.state.message.user
    const type = this.state.message.type
    //return <div className="Message">{this.state.message.body}</div>;
    if (type === 1) {
      return (
        <div className='Message'>
          <span
            className="avatar"
            style={{backgroundColor: this.state.message.user}}
          />
          <div className='message-content'>
            <div className="username">
              <p>ゲスト{userId}さん</p>
            </div>
            <div className="text">{this.state.message.body}</div>
          </div>
      </div>
      );
    };
    if (type === 2) {
      return (
        <div className="join">
          <p>{this.state.message.body}</p>
        </div>
      );
    }
  }
}

export default Message;