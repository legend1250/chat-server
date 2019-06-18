import { observable, action } from 'mobx'
import { JoinRoom } from '../views/home/components'

export interface roomTypes {
  roomId: string
}

class RootStore {
  constructor() {
    const ws = new WebSocket('ws://localhost:8080/ws')
    ws.onmessage = this.onListenMessage.bind(this)
    this.conn = ws
    this.loadingWS = false
  }

  @observable loadingWS: boolean = true
  @observable messages: Array<any> = []
  @observable roomInfo?: roomTypes = undefined
  conn: WebSocket

  onListenMessage(evt: any) {
    try {
      const obj = JSON.parse(evt.data)
      const { type } = obj
      switch (type) {
        // join room success
        case 2: {
          this.joinRoom(obj)
        }
        default: {
          this.messages.push(obj)
        }
      }
    } catch (error) {
      console.log('failed while listening message: ', error)
    }
  }

  @action joinRoom(data: any) {
    const { roomId } = data
    this.roomInfo = { roomId }
  }
}

export default RootStore
