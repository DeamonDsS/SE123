// rafce



import { userInfo } from './Store'
import { useRecoilValue } from 'recoil'
const Example = () => {
    const Best = useRecoilValue(userInfo)
    console.log(Best)
  return (

    <div>example{Best.detail}</div>
  )
}

export default Example

