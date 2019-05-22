import { Component, OnInit } from '@angular/core';
import {SocketService} from "../socket.service";

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.sass']
})
export class ChatComponent implements OnInit {

  msg
  nickName
  history:MessageBody[]=[]
  connectedList:Map<string,string> =new Map()
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

      switch (data.type) {
        case 1:
          this.history.push(data.body)
          break
        case 2:
          if(this.counter==1)
            this.myId=data.body.message
          else
              this.connectedList.set(data.body.message,data.body.nickname)
          break
        case 3:
          this.connectedList.delete(data.body.message)
          break
        case 4:
          this.connectedList.set(data.body.from,data.body.nickname)


      }

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
  }

  sendNickName(name){
    let socket = this.wsService.ws
    let body:MessageBody={
      from:this.myId,
      to:"",
      message:"",
      nickname:name
    }
    socket.send(JSON.stringify(body));
  }

}

export interface MessageBody {
  from
  to
  message
  nickname?
}

export interface SocketMessage {
  type
  body:MessageBody
}
