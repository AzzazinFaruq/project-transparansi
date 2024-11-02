<template>
  <div>
    <v-card class="pa-5">
      <div class="d-flex justify-space-between align-center mb-5">
        <div>
          <p class="text-h5 font-weight-bold">Tambah Program</p>
          <p class="text-subtitle-1 text-grey">Tambah program baru</p>
        </div>
        <v-btn 
          color="error" 
          @click="back()"
          prepend-icon="mdi-arrow-left"
        >
          Kembali
        </v-btn>
      </div>

      <v-divider class="mb-5"></v-divider>

      <v-row>
        <v-col cols="12" md="8">
          <v-form>
            <!-- Informasi Dasar -->
            <div class="mb-5">
              <label class="label-form mb-2 d-block">Nama Program</label>
              <v-text-field
                v-model="user.nama_program"
                placeholder="Masukkan nama program"
                variant="outlined"
                density="compact"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Institusi</label>
              <v-select
                v-model="user.institusi_id"
                :items="institusi"
                item-title="nama_institusi"
                item-value="id"
                placeholder="Pilih institusi"
                variant="outlined"
                density="compact"
              ></v-select>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Deskripsi</label>
              <v-textarea
                v-model="user.deskripsi"
                placeholder="Masukkan deskripsi program"
                variant="outlined"
                density="compact"
                rows="3"
              ></v-textarea>
            </div>

            <!-- Detail Anggaran -->
            <h2 class="text-h6 font-weight-bold mb-3">Detail Anggaran</h2>
            <v-divider class="mb-5"></v-divider>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Jenis Anggaran</label>
              <v-select
                v-model="user.jenis_anggaran_id"
                :items="JenisAnggaran"
                item-title="jenis"
                item-value="Id"
                placeholder="Pilih jenis anggaran"
                variant="outlined"
                density="compact"
              ></v-select>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Jumlah Anggaran</label>
              <v-text-field
                v-model="user.jumlah_anggaran"
                placeholder="Masukkan jumlah anggaran"
                variant="outlined"
                density="compact"
                prefix="Rp"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Kategori Penggunaan</label>
              <v-select
                v-model="user.kategori_penggunaan_id"
                :items="KategoriPenggunaan"
                item-title="kategori"
                item-value="Id"
                placeholder="Pilih kategori penggunaan"
                variant="outlined"
                density="compact"
              ></v-select>
            </div>

            <!-- Detail Alamat -->
            <h2 class="text-h6 font-weight-bold mb-3">Detail Alamat</h2>
            <v-divider class="mb-5"></v-divider>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Dusun</label>
              <v-text-field
                v-model="user.dusun"
                placeholder="Masukkan nama dusun"
                variant="outlined"
                density="compact"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Desa</label>
              <v-autocomplete
                v-model="user.desa_id"
                :items="deslist"
                item-title="nama_desa"
                item-value="Id"
                placeholder="Pilih desa"
                variant="outlined"
                density="compact"
              ></v-autocomplete>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Kecamatan</label>
              <v-autocomplete
                v-model="user.kecamatan_id"
                :items="keclist"
                item-title="nama_kecamatan"
                item-value="Id"
                placeholder="Pilih kecamatan"
                variant="outlined"
                density="compact"
              ></v-autocomplete>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Kabupaten / Kota</label>
              <v-autocomplete
                v-model="user.kabupaten_id"
                :items="kablist"
                item-title="nama_kabupaten"
                item-value="Id"
                placeholder="Pilih kabupaten"
                variant="outlined"
                density="compact"
              ></v-autocomplete>
            </div>

            <div class="d-flex justify-start">
              <v-btn variant="tonal" style="background-color:#387144;color: white;text-transform: none;" @click="create()">
                Simpan
              </v-btn>
            </div>
          </v-form>
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>
<script>
import axios from 'axios';
export default{
  data() {
    return {
      institusi:[],
      kablist:[],
      keclist:[],
      deslist:[],
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
        user_id:'',
        kecamatan_id:'',
        kabupaten_id:'',
        foto_before:'',
        foto_progress:'',
        foto_after:''


      }
    }
  },
  mounted() {
    this.userId();
    this.listinstitusi();
    this.listJenisAnggaran();
    this.listKategoriPenggunaan();
    this.listdaerah();
  },
  methods: {
    listdaerah(){
      axios.all([
    axios.get('/api/index-kabupaten'),  // Endpoint untuk kabupaten
    axios.get('/api/index-kecamatan'),  // Endpoint untuk kecamatan
    axios.get('/api/index-desa')        // Endpoint untuk desa
  ])
  .then(axios.spread((kabupatenRes, kecamatanRes, desaRes) => {
    this.kablist = kabupatenRes.data.data;
    this.keclist = kecamatanRes.data.data;
    this.deslist = desaRes.data.data;

  }))
  .catch(error => {
    console.error("Error fetching data:", error);
  });
    },
    userId(){
      axios.get("api/user")
      .then(res=>{
        this.user.user_id = res.data.data.Id
      })
    },
    listinstitusi(){
      axios.get("/api/institusi")
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
    create(){
      console.log(this.user);
    axios.post(`/api/program/pengajuan`, this.user)
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

