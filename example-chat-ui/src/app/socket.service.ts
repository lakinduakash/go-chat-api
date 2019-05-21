import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SocketService {

  private readonly _ws:WebSocket
  constructor() {
    this._ws = new WebSocket("ws://localhost:8080/ws")
  }

  get ws(){
    return this._ws
  }


}
