import { observable, action } from 'mobx'

class RootStore {
  constructor() {
    this.initWebsocket()
  }

  @observable loadingWS: boolean = true
  conn: any

  @action initWebsocket() {
    if ((window as any)['WebSocket']) {
      this.conn = new WebSocket('ws://localhost:8080/ws')
    }
    this.loadingWS = false
  }
}

export default RootStore
