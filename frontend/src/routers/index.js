import { createRouter , createWebHashHistory} from "vue-router"
import Home from "@/components/Home.vue"
import Login from "@/components/Login.vue"
import Player from "@/components/Player.vue"
import Team from "@/components/Team.vue"
import TeamPlayer from "@/components/TeamPlayer.vue"
import EditPlayer from "@/components/EditPlayer.vue"
import EditTeam from "@/components/EditTeam.vue"
import Signup from "@/components/Signup.vue"
import UserList from "@/components/UserList.vue"
const routes = [
  {
    path:"/home",
    component:Home,
    name:"home",
    children:[
      {path:"/login",component:Login,name:"login"},
      {path:"/player",component:Player,name:"player"},
      {path:"/team",component:Team,name:"team",children:[{path:"/teamPlayer",component:TeamPlayer}]},
      {path:"/editPlayer",component:EditPlayer,name:"editPlayer"},
      {path:"/editTeam",component:EditTeam,name:"editTeam"},
      {path:"/signup",component:Signup,name:"signup"},
      {path:"/userList",component:UserList,name:"userList"}
    ]
  },
  {
    path:"/",
    redirect:"/home"
  }
]

const router = createRouter({
  history:createWebHashHistory(),
  routes,
})

export default router