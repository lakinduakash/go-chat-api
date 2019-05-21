import { Component, OnInit } from '@angular/core';
import {SocketService} from "../socket.service";

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.sass']
})
export class ChatComponent implements OnInit {

  msg
  history:MessageEvent[]=[]

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
      console.log(msg);
      this.history.push(msg)
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
    console.log("sending msg: ", msg);
    socket.send(msg);
  };

}
