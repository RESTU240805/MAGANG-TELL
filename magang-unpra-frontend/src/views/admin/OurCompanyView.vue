<template>
  <div class="min-h-screen bg-gray-100 flex">

    <!-- ─── Sidebar ─── -->
    <aside class="w-64 bg-gray-950 text-white flex flex-col fixed h-full z-40">
      <div class="p-6 border-b border-gray-800">
        <div class="flex items-center gap-3">
          <div class="w-9 h-9 bg-green-600 rounded-xl flex items-center justify-center font-black text-sm">T</div>
          <div>
            <p class="font-black text-sm">TELPP</p>
            <p class="text-xs text-gray-500">Management Profile</p>
          </div>
        </div>
      </div>
      <nav class="flex-1 p-4 space-y-1">
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-2 px-2">OVERVIEW</p>
        <RouterLink to="/admin/dashboard"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📊 Dashboard
        </RouterLink>
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-4 px-2">CONTENT ENGINE</p>
        <RouterLink to="/admin/news"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📰 Corporate News
        </RouterLink>
        <RouterLink to="/admin/slider"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🖼️ Product Slider
        </RouterLink>
        <RouterLink to="/admin/product-page"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📝 Product Page
        </RouterLink>
        <RouterLink to="/admin/about"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🏢 About Section
        </RouterLink>
        <RouterLink to="/admin/team-members"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          👥 Our Team
        </RouterLink>
        <RouterLink to="/admin/our-company"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
          🏛️ Our Company
        </RouterLink>
        <RouterLink to="/admin/menus"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📋 Menu
        </RouterLink>
      </nav>
      <div class="p-4 border-t border-gray-800">
        <button @click="logout" class="text-sm text-gray-400 hover:text-white transition px-2">→ Logout</button>
      </div>
    </aside>

    <!-- ─── Main Content ─── -->
    <main class="flex-1 ml-64 p-10">

      <!-- Header -->
      <div class="flex justify-between items-center mb-8">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">Our Company</h1>
        </div>
      </div>

      <!-- Alert -->
      <div v-if="alert.show"
        :class="['mb-4 px-4 py-3 rounded-xl text-sm font-medium flex items-center gap-2',
          alert.type === 'success' ? 'bg-green-50 text-green-700 border border-green-200' : 'bg-red-50 text-red-700 border border-red-200']">
        {{ alert.type === 'success' ? '✅' : '❌' }} {{ alert.message }}
      </div>

      <!-- ════════════════════════════════════════════════════════
           SECTION: COMPANY PROFILE (Hero)
           ════════════════════════════════════════════════════════ -->
      <section class="bg-white rounded-2xl p-8 border border-gray-100 mb-8">
        <div class="flex items-center gap-3 mb-6">
          <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center text-blue-600 font-bold text-sm">P</div>
          <div>
            <h2 class="font-bold text-gray-900">Company Profile</h2>
            <p class="text-xs text-gray-400">Teks hero & deskripsi perusahaan (ditampilkan di bagian atas halaman Our Company)</p>
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Judul Perusahaan</label>
            <input v-model="profile.title" placeholder="PT Tanjungenim Lestari Pulp and Paper"
              class="w-full px-4 py-2.5 border border-gray-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Deskripsi (Hero)</label>
            <RichTextEditor v-model="profile.content" placeholder="Deskripsi perusahaan yang tampil di hero section..." />
          </div>
          <div class="grid grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Background Hero</label>
              <div class="flex items-center gap-4">
                <div v-if="profile.hero_image" class="w-28 h-20 rounded-xl overflow-hidden bg-gray-100 border border-gray-200 shrink-0 relative group">
                  <img :src="getImageUrl(profile.hero_image)" class="w-full h-full object-cover" />
                  <button @click="profile.hero_image = ''" type="button"
                    class="absolute top-1 right-1 w-5 h-5 bg-red-500 text-white rounded-full text-xs flex items-center justify-center opacity-0 group-hover:opacity-100 transition hover:bg-red-600">×</button>
                </div>
                <div class="flex-1">
                  <input type="file" accept="image/*" @change="onHeroImageChange" ref="heroInput" class="hidden" />
                  <div class="flex gap-2">
                    <button @click="$refs.heroInput.click()"
                      class="flex-1 border border-dashed border-gray-300 rounded-xl px-4 py-3 text-sm text-gray-500 hover:border-green-500 hover:text-green-600 transition text-left">
                      {{ uploading.hero ? 'Mengupload...' : 'Pilih gambar hero' }}
                    </button>
                    <button v-if="profile.hero_image" @click="profile.hero_image = ''" type="button"
                      class="px-3 py-3 text-xs font-medium text-red-600 bg-red-50 hover:bg-red-100 rounded-xl transition shrink-0">Hapus</button>
                  </div>
                  <p class="text-xs text-gray-400 mt-1">JPG, PNG. Maks 5MB. Jangan lupa klik "Simpan Profile" setelah mengganti.</p>
                </div>
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Background Creed</label>
              <div class="flex items-center gap-4">
                <div v-if="profile.creed_bg_image" class="w-28 h-20 rounded-xl overflow-hidden bg-gray-100 border border-gray-200 shrink-0 relative group">
                  <img :src="getImageUrl(profile.creed_bg_image)" class="w-full h-full object-cover" />
                  <button @click="profile.creed_bg_image = ''" type="button"
                    class="absolute top-1 right-1 w-5 h-5 bg-red-500 text-white rounded-full text-xs flex items-center justify-center opacity-0 group-hover:opacity-100 transition hover:bg-red-600">×</button>
                </div>
                <div class="flex-1">
                  <input type="file" accept="image/*" @change="onCreedBgImageChange" ref="creedBgInput" class="hidden" />
                  <div class="flex gap-2">
                    <button @click="$refs.creedBgInput.click()"
                      class="flex-1 border border-dashed border-gray-300 rounded-xl px-4 py-3 text-sm text-gray-500 hover:border-green-500 hover:text-green-600 transition text-left">
                      {{ uploading.creedBg ? 'Mengupload...' : 'Pilih gambar creed' }}
                    </button>
                    <button v-if="profile.creed_bg_image" @click="profile.creed_bg_image = ''" type="button"
                      class="px-3 py-3 text-xs font-medium text-red-600 bg-red-50 hover:bg-red-100 rounded-xl transition shrink-0">Hapus</button>
                  </div>
                  <p class="text-xs text-gray-400 mt-1">JPG, PNG. Maks 5MB. Jangan lupa klik "Simpan Profile" setelah mengganti.</p>
                </div>
              </div>
            </div>
          </div>
          <div class="flex justify-end pt-2">
            <button @click="saveProfile"
              class="bg-green-600 text-white px-6 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition flex items-center gap-2">
              💾 Simpan Profile
            </button>
          </div>
        </div>
      </section>

      <!-- ════════════════════════════════════════════════════════
           SECTION: COMPANY CREED (Kanji + Values)
           ════════════════════════════════════════════════════════ -->
      <section class="bg-white rounded-2xl p-8 border border-gray-100 mb-8">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 bg-amber-100 rounded-lg flex items-center justify-center text-amber-600 font-bold text-sm">C</div>
            <div>
              <h2 class="font-bold text-gray-900">Company Creed</h2>
              <p class="text-xs text-gray-400">Nilai perusahaan (kanji, roma, tagline, deskripsi) — muncul di bagian creed & values cards</p>
            </div>
          </div>
          <button @click="openCreedForm()"
            class="bg-green-600 text-white px-4 py-2 rounded-xl text-sm font-semibold hover:bg-green-700 transition flex items-center gap-1.5">
            <span class="text-lg leading-none">+</span> Tambah
          </button>
        </div>

        <!-- Loading -->
        <div v-if="loading" class="flex items-center justify-center py-10 gap-2 text-gray-400 text-sm">
          <div class="w-4 h-4 border-2 border-gray-200 border-t-green-500 rounded-full animate-spin"></div>
          Memuat data...
        </div>

        <!-- Daftar Creed -->
        <div v-else-if="creeds.length > 0" class="space-y-3">
          <div v-for="(creed, idx) in creeds" :key="creed.id"
            class="flex items-start justify-between p-4 bg-gray-50 rounded-xl border border-gray-100">
            <div class="flex-1 min-w-0">
              <div class="flex items-center gap-2 mb-1">
                <span class="text-xs font-mono text-gray-400 bg-gray-200 px-2 py-0.5 rounded">{{ idx }}</span>
                <span class="font-semibold text-gray-900">{{ creed.title_en }}</span>
                <span v-if="creed.title_jp" class="text-lg font-serif text-green-700">{{ creed.title_jp }}</span>
                <span v-if="creed.roma" class="text-xs text-gray-500">({{ creed.roma }})</span>
                <span :class="['text-xs px-2 py-0.5 rounded-full font-medium', creed.is_active ? 'bg-green-100 text-green-700' : 'bg-gray-200 text-gray-500']">
                  {{ creed.is_active ? 'Aktif' : 'Nonaktif' }}
                </span>
              </div>
              <p v-if="creed.tagline" class="text-sm text-green-600 italic">{{ creed.tagline }}</p>
              <p v-if="creed.description" class="text-sm text-gray-500 mt-1 line-clamp-2">{{ creed.description }}</p>
            </div>
            <div class="flex gap-2 ml-4 shrink-0">
              <button @click="openCreedForm(creed)"
                class="px-3 py-1.5 text-xs font-medium text-blue-600 bg-blue-50 hover:bg-blue-100 rounded-lg transition">Edit</button>
              <button @click="deleteCreed(creed.id)"
                class="px-3 py-1.5 text-xs font-medium text-red-600 bg-red-50 hover:bg-red-100 rounded-lg transition">Hapus</button>
            </div>
          </div>
        </div>

        <!-- Empty -->
        <div v-else class="text-center py-10 text-gray-400 text-sm">
          Belum ada creed. Klik "Tambah" untuk menambahkan nilai perusahaan.
        </div>

        <!-- Form Creed -->
        <div v-if="showCreedForm"
          class="mt-6 p-6 bg-gray-50 rounded-xl border border-gray-200">
          <h3 class="font-semibold text-gray-900 mb-4">{{ creedForm.id ? 'Edit Creed' : 'Tambah Creed Baru' }}</h3>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-xs font-medium text-gray-600 mb-1">Judul EN *</label>
              <input v-model="creedForm.title_en" placeholder="Fairness"
                class="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
            <div>
              <label class="block text-xs font-medium text-gray-600 mb-1">Kanji (JP)</label>
              <input v-model="creedForm.title_jp" placeholder="和"
                class="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
            <div>
              <label class="block text-xs font-medium text-gray-600 mb-1">Roma (pengucapan)</label>
              <input v-model="creedForm.roma" placeholder="WA"
                class="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
            <div>
              <label class="block text-xs font-medium text-gray-600 mb-1">Urutan tampil</label>
              <input v-model.number="creedForm.sort_order" type="number" placeholder="0"
                class="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
            <div class="col-span-2">
              <label class="block text-xs font-medium text-gray-600 mb-1">Tagline</label>
              <input v-model="creedForm.tagline" placeholder="To be fair and decent."
                class="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
            <div class="col-span-2">
              <label class="block text-xs font-medium text-gray-600 mb-1">Deskripsi</label>
              <textarea v-model="creedForm.description" rows="3" placeholder="Deskripsi nilai perusahaan..."
                class="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500"></textarea>
            </div>
            <div class="col-span-2">
              <label class="flex items-center gap-2 text-sm text-gray-700">
                <input type="checkbox" v-model="creedForm.is_active"
                  class="rounded border-gray-300 text-green-600 focus:ring-green-500" />
                Aktif (tampil di halaman publik)
              </label>
            </div>
          </div>
          <div class="flex gap-2 mt-4">
            <button @click="saveCreed"
              class="bg-green-600 text-white px-4 py-2 rounded-lg text-sm font-semibold hover:bg-green-700 transition">Simpan</button>
            <button @click="showCreedForm = false"
              class="bg-gray-200 text-gray-700 px-4 py-2 rounded-lg text-sm font-semibold hover:bg-gray-300 transition">Batal</button>
          </div>
        </div>
      </section>

      <!-- ════════════════════════════════════════════════════════
           SECTION: COMPANY DOCUMENT
           ════════════════════════════════════════════════════════ -->
      <section class="bg-white rounded-2xl p-8 border border-gray-100 mb-8">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center text-purple-600 font-bold text-sm">D</div>
            <div>
              <h2 class="font-bold text-gray-900">Document Preview</h2>
              <p class="text-xs text-gray-400">Dokumen PDF yang ditampilkan di bagian document preview</p>
            </div>
          </div>
          <button @click="openDocForm()"
            class="bg-green-600 text-white px-4 py-2 rounded-xl text-sm font-semibold hover:bg-green-700 transition flex items-center gap-1.5">
            <span class="text-lg leading-none">+</span> Tambah
          </button>
        </div>

        <!-- Daftar Dokumen -->
        <div v-if="documents.length > 0" class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead>
              <tr class="border-b border-gray-100">
                <th class="text-left py-3 px-2 font-semibold text-gray-500">Judul</th>
                <th class="text-left py-3 px-2 font-semibold text-gray-500">Kategori</th>
                <th class="text-left py-3 px-2 font-semibold text-gray-500">Tanggal</th>
                <th class="text-left py-3 px-2 font-semibold text-gray-500">File</th>
                <th class="text-right py-3 px-2 font-semibold text-gray-500">Aksi</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="doc in documents" :key="doc.id" class="border-b border-gray-50 hover:bg-gray-50">
                <td class="py-3 px-2 font-medium text-gray-900">{{ doc.title }}</td>
                <td class="py-3 px-2 text-gray-500">{{ doc.category }}</td>
                <td class="py-3 px-2 text-gray-500">{{ doc.doc_date }}</td>
                <td class="py-3 px-2 text-gray-500">{{ doc.file_type }} · {{ doc.file_size }}</td>
                <td class="py-3 px-2 text-right">
                  <button @click="openDocForm(doc)"
                    class="px-3 py-1.5 text-xs font-medium text-blue-600 bg-blue-50 hover:bg-blue-100 rounded-lg transition mr-1">Edit</button>
                  <button @click="deleteDocument(doc.id)"
                    class="px-3 py-1.5 text-xs font-medium text-red-600 bg-red-50 hover:bg-red-100 rounded-lg transition">Hapus</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <div v-else class="text-center py-10 text-gray-400 text-sm">
          Belum ada dokumen. Klik "Tambah" untuk menambahkan dokumen.
        </div>

        <!-- Form Dokumen -->
        <div v-if="showDocForm"
          class="mt-6 p-6 bg-gray-50 rounded-xl border border-gray-200">
          <h3 class="font-semibold text-gray-900 mb-4">{{ docForm.id ? 'Edit Dokumen' : 'Tambah Dokumen Baru' }}</h3>
          <div class="grid grid-cols-2 gap-4">
            <div class="col-span-2">
              <label class="block text-xs font-medium text-gray-600 mb-1">Judul</label>
              <input v-model="docForm.title" placeholder="PEFC Chain of Custody Commitment Statement"
                class="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
            <div>
              <label class="block text-xs font-medium text-gray-600 mb-1">Kategori</label>
              <input v-model="docForm.category" placeholder="DOCUMENT PREVIEW"
                class="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
            <div>
              <label class="block text-xs font-medium text-gray-600 mb-1">Tanggal</label>
              <input v-model="docForm.doc_date" placeholder="April 2023"
                class="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
            <div class="col-span-2">
              <label class="block text-xs font-medium text-gray-600 mb-1">Deskripsi</label>
              <textarea v-model="docForm.description" rows="3" placeholder="Deskripsi dokumen..."
                class="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500"></textarea>
            </div>
            <div class="col-span-2">
              <label class="block text-xs font-medium text-gray-600 mb-1">File PDF</label>
              <input type="file" accept=".pdf" @change="onFileChange"
                class="block w-full text-sm text-gray-500 file:mr-3 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-medium file:bg-green-50 file:text-green-700 hover:file:bg-green-100" />
              <p v-if="docForm.file_url" class="text-xs text-gray-400 mt-1">File saat ini: {{ docForm.file_url }}</p>
            </div>
          </div>
          <div class="flex gap-2 mt-4">
            <button @click="saveDocument"
              class="bg-green-600 text-white px-4 py-2 rounded-lg text-sm font-semibold hover:bg-green-700 transition">Simpan</button>
            <button @click="showDocForm = false"
              class="bg-gray-200 text-gray-700 px-4 py-2 rounded-lg text-sm font-semibold hover:bg-gray-300 transition">Batal</button>
          </div>
        </div>
      </section>

    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/services/api'
import RichTextEditor from '../../components/RichTextEditor.vue'

const BASE_URL = 'http://localhost:8080'
const router = useRouter()

const loading = ref(false)
const alert = ref({ show: false, type: 'success', message: '' })

const uploading = ref({ hero: false, creedBg: false })
const heroInput = ref(null)
const creedBgInput = ref(null)

const profile = ref({ title: '', content: '', hero_image: '', creed_bg_image: '' })
const creeds = ref([])
const documents = ref([])

const showCreedForm = ref(false)
const creedForm = ref({
  id: null,
  title_jp: '',
  title_en: '',
  roma: '',
  tagline: '',
  description: '',
  sort_order: 0,
  is_active: true,
})

const showDocForm = ref(false)
const docForm = ref({ id: null, title: '', category: '', doc_date: '', description: '', file_url: '' })
const selectedFile = ref(null)

onMounted(() => {
  const token = localStorage.getItem('token')
  if (!token) { router.push('/admin/login'); return }
  fetchProfile()
  fetchCreeds()
  fetchDocuments()
})

function showAlert(type, message) {
  alert.value = { show: true, type, message }
  setTimeout(() => { alert.value.show = false }, 3000)
}

// ─── PROFILE ───
async function fetchProfile() {
  try {
    const res = await api.get('/admin/company-profile')
    const data = res.data?.data || res.data
    if (data) {
      profile.value = {
        title: data.title || '',
        content: data.content || '',
        hero_image: data.hero_image || '',
        creed_bg_image: data.creed_bg_image || '',
      }
    }
  } catch (_e) {
    // pakai default
  }
}
async function saveProfile() {
  try {
    const res = await api.put('/admin/company-profile', profile.value)
    const data = res.data?.data || res.data
    if (data) {
      profile.value = {
        title: data.title || '',
        content: data.content || '',
        hero_image: data.hero_image || '',
        creed_bg_image: data.creed_bg_image || '',
      }
    }
    showAlert('success', 'Profile berhasil disimpan')
  } catch (e) {
    showAlert('error', 'Gagal menyimpan: ' + (e.response?.data?.error || e.message))
  }
}

function getImageUrl(path) {
  if (!path) { return '' }
  if (path.startsWith('http')) { return path }
  return BASE_URL + '/' + path.replace(/^\//, '')
}

async function onHeroImageChange(e) {
  const file = e.target.files?.[0]
  if (!file) { return }
  uploading.value.hero = true
  try {
    const fd = new FormData()
    fd.append('image', file)
    const res = await api.post('/upload', fd, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
    profile.value.hero_image = 'uploads/' + (res.data?.filename || '')
    showAlert('success', 'Gambar hero berhasil diupload!')
  } catch (_e) {
    showAlert('error', 'Gagal upload gambar hero')
  } finally {
    uploading.value.hero = false
  }
}

async function onCreedBgImageChange(e) {
  const file = e.target.files?.[0]
  if (!file) { return }
  uploading.value.creedBg = true
  try {
    const fd = new FormData()
    fd.append('image', file)
    const res = await api.post('/upload', fd, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
    profile.value.creed_bg_image = 'uploads/' + (res.data?.filename || '')
    showAlert('success', 'Gambar creed berhasil diupload!')
  } catch (_e) {
    showAlert('error', 'Gagal upload gambar creed')
  } finally {
    uploading.value.creedBg = false
  }
}

// ─── CREED ───
async function fetchCreeds() {
  loading.value = true
  try {
    const res = await api.get('/admin/creeds')
    creeds.value = Array.isArray(res.data) ? res.data : (res.data.data || [])
  } catch (_e) {
    // pakai default
  } finally {
    loading.value = false
  }
}
function openCreedForm(creed) {
  if (creed) {
    creedForm.value = { ...creed }
  } else {
    creedForm.value = {
      id: null, title_jp: '', title_en: '', roma: '',
      tagline: '', description: '', sort_order: creeds.value.length, is_active: true,
    }
  }
  showCreedForm.value = true
}
async function saveCreed() {
  try {
    if (creedForm.value.id) {
      await api.put(`/admin/creeds/${creedForm.value.id}`, creedForm.value)
    } else {
      await api.post('/admin/creeds', creedForm.value)
    }
    showCreedForm.value = false
    await fetchCreeds()
    showAlert('success', 'Creed berhasil disimpan')
  } catch (e) {
    showAlert('error', 'Gagal menyimpan: ' + (e.response?.data?.error || e.message))
  }
}
async function deleteCreed(id) {
  if (!confirm('Hapus creed ini?')) { return }
  try {
    await api.delete(`/admin/creeds/${id}`)
    await fetchCreeds()
    showAlert('success', 'Creed berhasil dihapus')
  } catch (_e) {
    showAlert('error', 'Gagal menghapus creed')
  }
}

// ─── DOCUMENT ───
async function fetchDocuments() {
  try {
    const res = await api.get('/admin/company-documents')
    documents.value = Array.isArray(res.data) ? res.data : (res.data.data || [])
  } catch (_e) {
    // pakai default
  }
}
function openDocForm(doc) {
  if (doc) {
    docForm.value = { ...doc }
  } else {
    docForm.value = { id: null, title: '', category: '', doc_date: '', description: '', file_url: '' }
  }
  selectedFile.value = null
  showDocForm.value = true
}
function onFileChange(e) {
  selectedFile.value = e.target.files[0]
}
async function saveDocument() {
  try {
    const formData = new FormData()
    formData.append('title', docForm.value.title)
    formData.append('category', docForm.value.category)
    formData.append('doc_date', docForm.value.doc_date)
    formData.append('description', docForm.value.description)
    if (selectedFile.value) {
      formData.append('file', selectedFile.value)
    }
    if (docForm.value.id) {
      await api.put(`/admin/company-documents/${docForm.value.id}`, formData, {
        headers: { 'Content-Type': 'multipart/form-data' },
      })
    } else {
      await api.post('/admin/company-documents', formData, {
        headers: { 'Content-Type': 'multipart/form-data' },
      })
    }
    showDocForm.value = false
    await fetchDocuments()
    showAlert('success', 'Dokumen berhasil disimpan')
  } catch (e) {
    showAlert('error', 'Gagal menyimpan: ' + (e.response?.data?.error || e.message))
  }
}
async function deleteDocument(id) {
  if (!confirm('Hapus dokumen ini?')) { return }
  try {
    await api.delete(`/admin/company-documents/${id}`)
    await fetchDocuments()
    showAlert('success', 'Dokumen berhasil dihapus')
  } catch (_e) {
    showAlert('error', 'Gagal menghapus dokumen')
  }
}

function logout() {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}
</script>
