import { defineStore } from 'pinia'
import axios from 'axios'
import { useRouter } from 'vue-router'

export const authStore = defineStore('auth', {
  state: () => {
    return{
      user_id:'',
      role:''
    }
  },
  actions:{
    userRole(){
      axios.get("/api/user")
      .then(res=>{
        this.role = res.data.data.Role.role
        console.log(this.role)
      })
    },
    check(router, swal){
      var stat = false;
        axios.get('/api/user')
        .then((res)=>{
          stat = res.data.status;
          this.user_id = res.data.Id

        })
        .catch(err=>{
          if (err.response.data.status ==  false) {
            router.push("/home")
            localStorage.removeItem('Role')
          }
        })



    }
  },
})
