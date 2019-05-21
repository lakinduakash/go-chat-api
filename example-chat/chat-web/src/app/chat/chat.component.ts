import { Component, OnInit } from '@angular/core';
import {SocketService} from "../socket.service";

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.sass']
})
export class ChatComponent implements OnInit {

  msg
  history:MessageBody[]=[]
  connectedList:string[]=[]
  toUser:string

  counter=0
  myId:string

  constructor(public wsService:SocketService) { }

  ngOnInit() {
    this.connect()
  }

  connect (){
    let socket = this.wsService.ws
    console.log("Attempting Connection...");

    socket.onopen = () => {
      console.log("Successfully Connected");
    };

    socket.onmessage = msg => {
      console.log(msg)
      this.counter++
      let data =JSON.parse(msg.data) as SocketMessage

      if(data.type ==1)
        this.history.push(data.body)
      else if(data.type ==2) {
        if(this.counter==1)
          this.myId=data.body.message
        else
          this.connectedList.push(data.body.message)
      }
      else if(data.type ==3)
        this.connectedList=this.connectedList.filter(val=>val!= data.body.message)

    };

    socket.onclose = event => {
      console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
      console.log("Socket Error: ", error);
    };
  }

  sendMsg(msg) {
    let socket = this.wsService.ws
    let body:MessageBody={
      from:this.myId,
      to:this.toUser,
      message:msg
    }
    socket.send(JSON.stringify(body));
  };

}

export interface MessageBody {
  from
  to
  message
}

export interface SocketMessage {
  type
  body:MessageBody
}
