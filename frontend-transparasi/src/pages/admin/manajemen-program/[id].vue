<template>
  <v-container>
    <div class="" v-if="user.status == 'Menunggu' || user.status == 'Disetujui'">
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

        <div class="d-flex justify-start" v-if="user.status=='Menunggu'">
          <v-btn variant="tonal" style="background-color:#387144;color: white;text-transform: none;" @click="accept()">
              Terima
            </v-btn>
            <v-btn variant="tonal" class="ml-2" style="background-color:#BF3232;color: white;text-transform: none;" @click="rejectprogram()">
              Tolak
            </v-btn>
        </div>
        <div class="d-flex justify-start" v-if="user.status=='Disetujui'">
          <v-btn variant="tonal" style="background-color:#387144;color: white;text-transform: none;" @click="update()">
              Simpan
            </v-btn>
            <v-btn variant="tonal" class="ml-2" style="background-color:#BF3232;color: white;text-transform: none;" @click="back()">
              Kembali
            </v-btn>
        </div>
      </v-form>
    </div>
    <div class="" v-if="user.status=='Ditolak'">
      <label class="label-form">Nama Program</label>
      <p>{{ user.nama_program }}</p>
      <label class="label-form">Institusi</label>
      <p>{{ user.Institusi.nama_institusi }}</p>
      <label class="label-form">Deskripsi</label>
      <p>{{ user.deskripsi }}</p>
      <label class="label-form">Deskripsi</label>
      <p>{{ user.deskripsi }}</p>

      <h2 class="my-3">Detail Anggaran</h2>
      <VDivider/>

      <label class="label-form">Jenis Anggaran</label>
      <p>{{ user.JenisAnggaran.jenis }}</p>
      <label class="label-form">Jumlah Anggaran</label>
      <p>{{ user.jumlah_anggaran }}</p>
      <label class="label-form">Kategori Penggunaan</label>
      <p>{{ user.KategoriPenggunaan.kategori}}</p>

      <h2 class="my-3">Detail Alamat</h2>
      <VDivider/>

      <label class="label-form">Dusun</label>
      <p>{{ user.nama_program }}</p>
      <label class="label-form">Desa</label>
      <p>{{ user.desa_id }}</p>
      <label class="label-form">Kecamatan</label>
      <p>{{ user.kecamatan_id }}</p>
      <label class="label-form">Kabupaten</label>
      <p>{{ user.kabupaten_id  }}</p>


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
        jenis_anggaran_id:'',
        jumlah_anggaran:'',
        kategori_penggunaan_id:'',
        institusi_id:'',
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
    accept(){
    axios.get(`/api/program/accept/${this.$route.params.id}`)
    .then(res=>{
      this.$router.go(0)
    })
  },
  rejectprogram(){
    axios.get(`/api/program/reject/${this.$route.params.id}`)
    .then(this.$router.go(-1))
  },
  back(){
    this.$router.go(-1)
  }
  },

}
</script>

