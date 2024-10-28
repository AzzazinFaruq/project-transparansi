<template>
  <v-container>
    <div class="">
      <v-form>
        <label class="label-form">Nama Program</label>
        <v-text-field v-model="user.nama_program" variant="outlined"></v-text-field>
        <label class="label-form">Institusi</label>
        <v-select
            v-model="user.institusi_id"
            :items="institusi"
            item-title="nama_institusi"
            item-value="Id"
            variant="outlined"
          ></v-select>
        <label  class="label-form">Deskripsi</label>
        <v-textarea v-model="user.deskripsi" variant="outlined"></v-textarea>

        <h2 class="my-3">Detail Anggaran</h2>

        <label class="label-form">Jenis Anggaran</label>
        <v-select
            v-model="user.jenis_anggaran_id"
            :items="JenisAnggaran"
            item-title="jenis"
            item-value="Id"
            variant="outlined"
          ></v-select>
        <label class="label-form">Jumlah Anggaran</label>
        <v-text-field v-model="user.jumlah_anggaran" variant="outlined"></v-text-field>
        <label class="label-form">kategori Penggunaan</label>
        <v-select
            v-model="user.kategori_penggunaan_id"
            :items="KategoriPenggunaan"
            item-title="kategori"
            item-value="Id"
            variant="outlined"
          ></v-select>

        <h2 class="my-3">Detail Alamat</h2>

        <label class="label-form">Dusun</label>
        <v-text-field v-model="user.dusun" variant="outlined"></v-text-field>
        <label class="label-form">Desa</label>
        <v-text-field v-model="user.desa_id" variant="outlined"></v-text-field>
        <label class="label-form">Kecamatan</label>
        <v-text-field v-model="user.kecamatan_id" variant="outlined"></v-text-field>
        <label class="label-form">Kabupaten / Kota</label>
        <v-text-field v-model="user.kabupaten_id" variant="outlined"></v-text-field>

        <div class="" v-if="user.status == 'Disetujui'">
          <h2 class="my-3">Detail Dokumentasi</h2>

          <label class="label-form">Before</label>
          <v-text-field v-model="user.foto_before" variant="outlined"></v-text-field>
          <label class="label-form">Proses</label>
          <v-text-field v-model="user.foto_progress" variant="outlined"></v-text-field>
          <label class="label-form">After</label>
          <v-text-field v-model="user.foto_after" variant="outlined"></v-text-field>
        </div>
      </v-form>
    </div>
  </v-container>
</template>
<script>
import axios from 'axios';
export default{
  data() {
    return {
      institusi:[],
      JenisAnggaran:[],
      KategoriPenggunaan:[],
      user:{
        nama_program:'',
        deskripsi: '',
        institusi_id:'',
        jenis_anggaran_id:'',
        jumlah_anggaran:'',
        kategori_penggunaan_id:'',
        dusun:'',
        desa_id:'',
        kecamatan_id:'',
        kabupaten_id:'',
        foto_before:'',
        foto_progress:'',
        foto_after:''


      }
    }
  },
  mounted() {
    this.getProgram();
    this.listJenisAnggaran();
    this.listKategoriPenggunaan();
    this.listinstitusi();
  },
  methods: {
    listinstitusi(){
      axios.get("/api/index-institusi")
      .then(res=>{
        this.institusi = res.data.data
      })
    },
    listJenisAnggaran(){
      axios.get("/api/index-jenis-anggaran")
      .then(res=>{
        this.JenisAnggaran = res.data.data
      })
    },
    listKategoriPenggunaan(){
      axios.get("/api/index-kategori-penggunaan")
      .then(res=>{
        this.KategoriPenggunaan = res.data.data
      })
    },
    getProgram(){
      axios.get(`/api/program/${this.$route.params.id}`)
      .then(res=>{
        console.log(res.data.data)
        this.user = res.data.data
      })
    },
    update(){
    axios.put(`/api/user/edituser/${this.$route.params.id}`, this.user)
    .then(res=>{
      this.$router.go(-1)
    })
  },
  back(){
    this.$router.go(-1)
  }
  },

}
</script>

