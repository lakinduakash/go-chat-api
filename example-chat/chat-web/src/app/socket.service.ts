import { Injectable } from '@angular/core';
import {environment} from "../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class SocketService {

  private readonly _ws:WebSocket
  constructor() {
    this._ws = new WebSocket(`ws://${environment.chat_sever}/ws`)
  }

  get ws(){
    return this._ws
  }


}
