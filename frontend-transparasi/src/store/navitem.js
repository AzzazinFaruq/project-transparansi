import { defineStore } from 'pinia'
import axios from 'axios'
import router from '@/router'


export const navitemstore = defineStore('navItem', {
  state: () => {
    return{
      navitemlist:[],
      stat:'',
      reload:false,
      itemAdmin: [
        {
          text: 'Dashboard',
          icon: 'mdi-home',
          route:'/admin/dashboard'
        },
        {
          text: 'Manajemen Program',
          icon: 'mdi-list-box',
          route:'/admin/manajemen-program'
        },
        {
          text: 'Manajemen Keluhan',
          icon: 'mdi-email',
          route:'/admin/keluhan'
        },
        {
          text: 'Manajemen Institusi',
          icon: 'mdi-office-building',
          route:'/admin/institusi'
        },
        {
          text: 'Manajemen Jabatan',
          icon: 'mdi-office-building',
          route:'/admin/jabatan'
        },


      ],
      itemAnggota: [
        {
          text: 'Dashboard',
          icon: 'mdi-home',
          route:'/dprd/dashboard'
        },
        {
          text: 'Program',
          icon: 'mdi-list-box',
          route:'/dprd/manajemen-program'
        },
        {
          text: 'Keluhan',
          icon: 'mdi-email',
          route:'/dprd/keluhan'
        },
        {
          text: 'History',
          icon: 'mdi-history',
          route:'/dprd/history'
        },

      ],
    }
  },
  actions:{
    reset() {
      this.navitemlist=[]
    },
    check(){
        axios.get('/api/user')
        .then((res)=>{
          this.stat=res.data.data.Role.role;
          console.log(this.stat)
          switch (this.stat) {
            case 'user':
              break;
            case 'dprd':
              this.navitemlist = this.itemAnggota;
              break;
            case 'admin':
              this.navitemlist = this.itemAdmin;
              break;
            default:
              break;
          }
          if (this.reload == true) {
            this.reload=false
            window.location.reload();
          }
        })
        .catch(err=>{
        })



    }
  },
})
