<template>
  <div class="app-wrapper"> 
    <div class="app-content pt-3 p-md-3 p-lg-4">
      <div class="container-xl">
  <div class="tab-content" id="orders-table-tab-content">
    <div class="tab-pane fade show active" id="orders-all" role="tabpanel" aria-labelledby="orders-all-tab">
      <div class="app-card app-card-orders-table shadow-sm mb-5">
        <div class="app-card-body">
          <div class="table-responsive">
            <table class="table app-table-hover mb-0 text-left">
              <thead>
                <tr>
                  <th class="cell"></th>
                  <th class="cell">Cn Name</th>
                  <th class="cell">En Name</th>
                  <th class="cell">Height</th>
                  <th class="cell">Weight</th>
                  <th class="cell">Nation</th>
                  <th class="cell">Age</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="player in (teamPlayer)" :key="player.pid">
                  <td class="cell">
                    <img style="height:40px" :src="('./img/player_img/' + player.imgpath)" alt="player_img">
                  </td>
                  <td class="cell">{{ player.cnname }}</td>
                  <td class="cell">{{ player.enname }}</td>
                  <td class="cell">{{ player.height }}</td>
                  <td class="cell"><span>{{ player.weight }}</span></td>
                  <td class="cell"><span>{{ player.nation }}</span></td>
                  <td class="cell"><span>{{ player.age }}</span></td>            
                </tr>
              </tbody>
            </table>
          </div><!--//table-responsive-->
              
        </div><!--//app-card-body-->		
      </div><!--//app-card-->
      
    </div><!--//tab-pane-->
  </div>
</div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
 export default{
  data(){
    return {
      teamPlayer:[],
    }
  },
  mounted(){
    axios.get("/api/player/teamPlayer",{
      params:{
        id:this.$route.query.tid
      }
    }).then(response => {
      this.teamPlayer = response.data.obj
    })
  },
  watch:{
    $route(){ axios.get("/api/player/teamPlayer",{
        params:{
          id:this.$route.query.tid
        }
      }).then(response => {
        this.teamPlayer = response.data.obj
      })
      
    }
  }
 }
</script>