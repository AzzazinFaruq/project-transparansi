import { defineStore } from 'pinia'
import axios from 'axios'

export const navitemstore = defineStore('navItem', {
  state: () => {
    return{
      navitemlist:[],
      stat:[],
      itemAdmin: [
        {
          text: 'Dashboard',
          icon: 'mdi-home',
          route:'/admin/dashboard'
        },
        {
          text: 'Manajemen Pengguna',
          icon: 'mdi-account',
          route:'/admin/manajemen-pengguna'
        },
        {
          text: 'Manajemen Program',
          icon: 'mdi-list-box',
          route:'/admin/manajemen-program'
        },
        {
          text: 'Keluhan',
          icon: 'mdi-list-box',
          route:'/admin/keluhan'
        },

      ],
    }
  },
  actions:{
    check(){

        axios.get('/api/user')
        .then((res)=>{
          this.stat=res.data.data.Role.role;
          console.log(this.stat)
          switch (this.stat) {
            case 'user':
              break;
            case 'anggota':
              break;
            case 'admin':
              this.navitemlist = this.itemAdmin;
              break;

            default:
              break;
          }
        })
        .catch(err=>{
          // if (err.response.data.status ==  false) {
          //   router.push("/login")
          //   swal({
          //     toast: "true",
          //     timer:2000,
          //     position:"top-end",
          //     icon: "error",
          //     title: "Error ",
          //     text:"Error Memuat data Navigasi",
          //     showConfirmButton :false
          //     });
          // }
        })



    }
  },
})