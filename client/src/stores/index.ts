import rootStore, { roomTypes } from './rootStore'

export { rootStore }
export interface StoresTypes {
  loadingWS: boolean
  messages: Array<any>
  conn: WebSocket
  roomInfo: roomTypes
}
