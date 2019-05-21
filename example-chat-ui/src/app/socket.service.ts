import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SocketService {

  private readonly _ws:WebSocket
  constructor() {
    this._ws = new WebSocket("ws://localhost:28960/ws")
  }

  get ws(){
    return this._ws
  }


}
