<template>
 <v-container>
  <div class="my-5">
    <v-btn variant="text" @click="router.go(-1)" prepend-icon="mdi-chevron-left">Kembali</v-btn>
  </div>
  <div class="mt-10">
    <v-row>
      <v-col cols="12" md="3">
        <div class="d-flex justify-center">
          <v-img v-if="user.foto_profil == ''" src="@/assets/profile.png" width="80%" height="auto"></v-img>
          <v-img :src="`${getImageUrl(user.foto_profil)}`" width="80%" height="auto"></v-img>
        </div>
      </v-col>
      <v-col cols="12" md="9">
        <v-row>
          <v-col cols="12" sm="6">
          <div class="mb-10">
          <label for="" class="label-form">Nama</label>
            <p>{{ user.username }}</p>
          </div>
          <div class="mb-10" v-if="user.role_id == 1 || user.role_id == 2">
          <label for="" class="label-form">Jabatan</label>
            <p>{{ jabatan }}</p>
          </div>
        <div class="mb-10">
          <label for="" class="label-form">No. HP</label>
            <p>{{ user.no_hp }}</p>
        </div>
          </v-col>
          <v-col cols="12" sm="6">
            <div class="mb-10">
          <label for="" class="label-form">Email</label>
          <p>{{ user.email }}</p>
        </div>
        <div class="">
          <label for="" class="label-form">Alamat</label>
          <p>{{ user.alamat }}</p>
        </div>
          </v-col>
        </v-row>
        <div class="mt-10">
           <v-btn color="#387144" prepend-icon="mdi-pencil" @click="router.push(`/profile/${user.Id}`)" style="text-transform: none;">Ubah Profil</v-btn>
  </div>     
      </v-col>
    </v-row>
  </div>
  
 </v-container>
</template>
<script>
import { useRouter } from 'vue-router';
import axios from 'axios';
import { getImageUrl } from '@/config/foto';

export default {
  setup() {
    const router = useRouter();
    return { router };
  },
  data:()=>({
    getImageUrl,
    user:[],
    jabatan:''
  }),
  mounted(){
    this.Getuser();
  },
  methods:{
    Getuser(){
      axios.get('/api/user')
      .then(res=>{
        this.user = res.data.data
        this.jabatan =res.data.data.Jabatan.jabatan
      })
    }
  }
};
</script>
