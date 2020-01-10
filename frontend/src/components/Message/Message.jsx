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
    //return <div className="Message">{this.state.message.body}</div>;
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
  }
}

export default Message;