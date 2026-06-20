<template>
  <div class="page">

    <!-- HERO -->
    <section class="hero">
      <div class="hero-bg">
        <img src="/images/lokasi pabrik.jpeg" alt="Our Team" class="hero-img" />
        <div class="hero-overlay"></div>
      </div>
      <div class="hero-content">
        <p class="eyebrow in">OUR COMPANY</p>
        <h1 class="in">
          <span class="line">OUR TEAM</span>
        </h1>
        <div class="accent-line in"></div>
        <p class="hero-p in">
          Meet the talented people who helped make PT TELPP what it is today.
        </p>
      </div>
    </section>

    <!-- TEAM LIST -->
    <section class="team-section">
      <div class="section-inner">
        <h2 class="section-title">Dewan Direksi</h2>

        <div v-if="loading" class="flex items-center justify-center py-24">
          <div class="w-6 h-6 border-2 border-emerald-600 border-t-transparent rounded-full animate-spin"></div>
        </div>

        <div v-if="featuredMembers.length" class="featured-list">
          <div v-for="featured in featuredMembers" :key="featured.ID" class="featured-card">
            <div class="featured-photo">
              <span class="featured-photo-placeholder">👤</span>
              <img v-if="featured.photo_path" :src="getImageUrl(featured.photo_path)" :alt="featured.name" />
            </div>
            <div class="featured-info">
              <span class="featured-badge">Petinggi Perusahaan</span>
              <h3 class="featured-name">{{ featured.name }}</h3>
              <p class="featured-pos">{{ featured.position }}</p>
              <p class="featured-desc">{{ featured.description }}</p>
              <button @click="openDetail(featured)" class="featured-btn">Selengkapnya →</button>
            </div>
          </div>
        </div>

        <div v-if="regulars.length" class="team-grid">
          <div
            v-for="(m, i) in regulars"
            :key="m.ID"
            class="team-card"
            :style="{ animationDelay: (i * 0.12) + 's' }"
          >
            <div class="team-photo">
              <div class="photo-placeholder">👤</div>
              <img v-if="m.photo_path" :src="getImageUrl(m.photo_path)" :alt="m.name" class="team-img" />
            </div>
            <div class="team-info">
              <h3 class="team-name">{{ m.name }}</h3>
              <p class="team-position">{{ m.position }}</p>
              <p v-if="m.description" class="team-desc">{{ truncate(m.description, 120) }}</p>
              <button @click="openDetail(m)" class="read-more-btn">Selengkapnya →</button>
            </div>
          </div>
        </div>

        <div v-else class="empty-state">
          <p class="empty-icon">👥</p>
          <p class="empty-text">Belum ada data anggota tim.</p>
        </div>
      </div>
    </section>

    <!-- MODAL DETAIL ANGGOTA -->
    <Transition name="modal">
      <div v-if="showDetail" class="modal-overlay" @click.self="showDetail = false">
        <!-- MODAL: Featured / Petinggi (horizontal besar) -->
        <div v-if="selectedMember?.is_featured" class="modal-card modal-featured">
          <button @click="showDetail = false" class="modal-x">×</button>
          <div class="mf-photo">
            <span v-if="!selectedMember?.photo_path" class="mf-photo-placeholder">👤</span>
            <img v-if="selectedMember?.photo_path" :src="getImageUrl(selectedMember.photo_path)" :alt="selectedMember.name" />
            <div class="mf-badge">PETINGGI</div>
          </div>
          <div class="mf-info">
            <h3 class="mf-name">{{ selectedMember?.name }}</h3>
            <p class="mf-pos">{{ selectedMember?.position }}</p>
            <div class="mf-divider"></div>
            <p class="mf-desc">{{ selectedMember?.description || 'Tidak ada deskripsi.' }}</p>
          </div>
        </div>

        <!-- MODAL: Regular -->
        <div v-else class="modal-card modal-regular">
          <button @click="showDetail = false" class="modal-x">×</button>
          <div class="mr-photo">
            <span v-if="!selectedMember?.photo_path" class="modal-photo-placeholder">👤</span>
            <img v-if="selectedMember?.photo_path" :src="getImageUrl(selectedMember.photo_path)" :alt="selectedMember.name" />
          </div>
          <div class="mr-info">
            <h3 class="mr-name">{{ selectedMember?.name }}</h3>
            <p class="mr-pos">{{ selectedMember?.position }}</p>
            <p class="mr-desc">{{ selectedMember?.description || 'Tidak ada deskripsi.' }}</p>
          </div>
        </div>
      </div>
    </Transition>

    <!-- STRUKTUR ORGANISASI (gambar bagan) -->
    <section class="orgchart-section">
      <div class="section-inner">
        <h2 class="section-title">Struktur Organisasi</h2>
        <div v-if="orgChartImage" class="orgchart-img-wrap">
          <img :src="getImageUrl(orgChartImage)" alt="Struktur Organisasi" class="orgchart-img" />
        </div>
        <div v-else class="empty-state">
          <p class="empty-icon">🏢</p>
          <p class="empty-text">Struktur organisasi belum tersedia.</p>
        </div>
      </div>
    </section>

    <footer class="site-footer">
      <div class="footer-container">
        <div class="footer-left-content"></div>
        <div class="footer-copyright">
          <p>Copyright 2026 PT TELPP. All right reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import api from '../services/api'

const BASE_URL = 'http://localhost:8080'
const members = ref([])
const loading = ref(true)
const showDetail = ref(false)
const selectedMember = ref(null)

const featuredMembers = computed(() => members.value.filter(m => m.is_featured))
const regulars = computed(() => members.value.filter(m => !m.is_featured))

const getImageUrl = (path) => {
  if (!path) {return ''}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const openDetail = (m) => {
  selectedMember.value = m
  showDetail.value = true
}

const truncate = (text, max = 120) => {
  if (!text || text.length <= max) {return text}
  return text.slice(0, max) + '...'
}

const orgChartImage = ref('')

onMounted(async () => {
  try {
    const res = await api.get('/team-members')
    members.value = res.data?.data || []
  } catch {
    members.value = []
  } finally {
    loading.value = false
  }
  try {
    const res = await api.get('/org-chart')
    orgChartImage.value = res.data?.data?.image_path || ''
  } catch (_e) { /* org-chart optional */ }
})
</script>

<style scoped>
* { box-sizing: border-box; margin: 0; padding: 0; }
.page {
  font-family: 'Segoe UI', system-ui, Arial, sans-serif;
  color: #1a1a1a;
  background: #fff;
  overflow-x: hidden;
}

/* ── HERO ── */
.hero {
  position: relative;
  min-height: 400px;
  display: flex;
  align-items: flex-end;
  overflow: hidden;
}
.hero-bg { position: absolute; inset: 0; z-index: 0; }
.hero-img {
  width: 100%; height: 100%; object-fit: cover;
  object-position: center 40%;
  display: block;
  transform: scale(1.06);
  animation: heroZoom 14s ease-out forwards;
}
@keyframes heroZoom { to { transform: scale(1); } }

.hero-overlay {
  position: absolute; inset: 0;
  background: linear-gradient(105deg, rgba(5,28,14,.88) 0%, rgba(5,28,14,.6) 42%, rgba(5,28,14,.18) 100%);
}

.hero-content {
  position: relative; z-index: 2;
  padding: 64px 72px 72px;
  max-width: 660px;
}

.eyebrow {
  font-size: .7rem; font-weight: 700; letter-spacing: .16em;
  color: #5ecb7a; margin-bottom: 18px;
  opacity: 0; transform: translateY(14px);
  transition: opacity .6s, transform .6s;
}
.eyebrow.in { opacity: 1; transform: none; }

h1 { margin: 0 0 22px; color: #fff; }
.line {
  display: block;
  font-size: 2.8rem; font-weight: 800; line-height: 1.18;
  opacity: 0; transform: translateY(22px) skewY(1.5deg);
  transition: opacity .7s cubic-bezier(.22,1,.36,1), transform .7s cubic-bezier(.22,1,.36,1);
}
.in .line { opacity: 1; transform: none; }

.accent-line {
  width: 0; height: 3px; background: #2d7a3d;
  border-radius: 2px; margin-bottom: 22px;
  transition: width .8s cubic-bezier(.22,1,.36,1);
}
.accent-line.in { width: 56px; }

.hero-p {
  font-size: .93rem; color: rgba(255,255,255,.8); line-height: 1.8;
  opacity: 0; transform: translateY(14px);
  transition: opacity .7s, transform .7s;
}
.hero-p.in { opacity: 1; transform: none; }

/* ── TEAM SECTION ── */
.team-section {
  padding: 72px 72px 88px;
  background: #fff;
}
.section-inner { max-width: 1200px; margin: 0 auto; }

.section-title {
  font-size: 1.8rem;
  font-weight: 800;
  color: #1a1a1a;
  margin-bottom: 40px;
}

/* ── FEATURED CARD (Horizontal Memanjang) ── */
.featured-list {
  display: flex;
  flex-direction: column;
  gap: 32px;
  margin-bottom: 48px;
}
.featured-card {
  display: flex;
  gap: 40px;
  background: linear-gradient(135deg, #fafcf8 0%, #f0f7ee 100%);
  border: 2px solid #e2e8f0;
  border-radius: 20px;
  padding: 0;
  margin-bottom: 48px;
  transition: box-shadow .3s, border-color .3s;
  position: relative;
  overflow: hidden;
  min-height: 300px;
}
.featured-card:hover {
  border-color: #cbd5e1;
  box-shadow: 0 12px 40px rgba(0,0,0,.08);
}

.featured-photo {
  width: 280px;
  min-height: 300px;
  flex-shrink: 0;
  position: relative;
  overflow: hidden;
}
.featured-photo img {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform .5s;
}
.featured-card:hover .featured-photo img { transform: scale(1.04); }
.featured-photo-placeholder {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 5rem;
  color: #bbb;
  background: #e8f0e4;
}
.featured-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 40px 40px 40px 0;
  min-width: 0;
}
.featured-badge {
  display: inline-block;
  width: fit-content;
  padding: 6px 16px;
  border-radius: 8px;
  font-size: .72rem;
  font-weight: 700;
  background: #0f172a;
  color: #fff;
  margin-bottom: 14px;
  text-transform: uppercase;
  letter-spacing: .06em;
}
.featured-name {
  font-size: 1.8rem;
  font-weight: 800;
  color: #0f172a;
  margin: 0 0 4px;
  letter-spacing: -.02em;
}
.featured-pos {
  font-size: 1.05rem;
  font-weight: 600;
  color: #2563eb;
  margin: 0 0 16px;
}
.featured-desc {
  font-size: .92rem;
  color: #475569;
  line-height: 1.8;
  margin: 0 0 20px;
  max-width: 600px;
}
.featured-btn {
  align-self: flex-start;
  padding: 10px 24px;
  border-radius: 10px;
  font-size: .85rem;
  font-weight: 600;
  background: #0f172a;
  color: #fff;
  border: none;
  cursor: pointer;
  transition: background .2s, transform .2s;
}
.featured-btn:hover { background: #1e293b; transform: translateY(-1px); }

/* ── GRID LAYOUT ── */
.team-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 36px 32px;
}

.team-card {
  display: flex;
  flex-direction: column;
  animation: fadeInUp .6s cubic-bezier(.22,1,.36,1) both;
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(24px); }
  to { opacity: 1; transform: none; }
}

.team-photo {
  width: 100%;
  aspect-ratio: 1 / 1;
  border-radius: 12px;
  overflow: hidden;
  background: #e8f0e4;
  position: relative;
  margin-bottom: 20px;
}
.team-photo .team-img {
  position: absolute;
  inset: 0;
  width: 100%; height: 100%; object-fit: cover;
  transition: transform .5s cubic-bezier(.22,1,.36,1);
}
.team-card:hover .team-photo .team-img { transform: scale(1.04); }

.photo-placeholder {
  position: absolute;
  inset: 0;
  width: 100%; height: 100%;
  display: flex; align-items: center; justify-content: center;
  font-size: 4rem; color: #bbb;
}

.team-info {
  text-align: left;
}
.team-name {
  font-size: 1.15rem;
  font-weight: 800;
  color: #1a1a1a;
  margin-bottom: 4px;
  letter-spacing: -.01em;
}
.team-position {
  font-size: .88rem;
  color: #2563eb;
  font-weight: 500;
  margin-bottom: 8px;
}
.team-desc {
  font-size: .85rem;
  color: #6b7280;
  line-height: 1.6;
  margin-bottom: 6px;
}
.read-more-btn {
  background: none;
  border: none;
  color: #2563eb;
  font-weight: 600;
  font-size: .82rem;
  cursor: pointer;
  padding: 0;
}
.read-more-btn:hover {
  text-decoration: underline;
}

/* ── MODAL ── */
.modal-enter-active, .modal-leave-active { transition: opacity .25s ease; }
.modal-enter-from, .modal-leave-to { opacity: 0; }

.modal-overlay {
  position: fixed;
  inset: 0;
  z-index: 999;
  background: rgba(0,0,0,.5);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}
.modal-card {
  background: #fff;
  border-radius: 16px;
  overflow: hidden;
  position: relative;
  box-shadow: 0 20px 60px rgba(0,0,0,.18);
  max-height: 86vh;
}
.modal-x {
  position: absolute;
  top: 12px; right: 12px;
  width: 30px; height: 30px;
  border-radius: 50%;
  border: none;
  background: #dc2626;
  color: #fff;
  font-size: 1.1rem;
  font-weight: 700;
  cursor: pointer;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background .2s;
  line-height: 1;
}
.modal-x:hover { background: #b91c1c; }

/* ── FEATURED MODAL (horizontal besar) ── */
.modal-featured {
  max-width: 820px;
  width: 100%;
  display: flex;
  height: 420px;
}
.mf-photo {
  width: 380px;
  flex-shrink: 0;
  background: #e8f0e4;
  position: relative;
  overflow: hidden;
}
.mf-photo img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.mf-photo-placeholder {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 5rem;
  color: #bbb;
}
.mf-badge {
  position: absolute;
  top: 16px; left: 16px;
  padding: 5px 14px;
  border-radius: 6px;
  font-size: .72rem;
  font-weight: 800;
  background: #fef3c7;
  color: #92400e;
  text-transform: uppercase;
  letter-spacing: .04em;
}
.mf-info {
  flex: 1;
  padding: 44px 40px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.mf-name {
  font-size: 1.8rem;
  font-weight: 800;
  color: #0f172a;
  margin: 0 0 4px;
}
.mf-pos {
  font-size: 1rem;
  font-weight: 600;
  color: #2563eb;
  margin: 0 0 16px;
}
.mf-divider {
  width: 50px;
  height: 3px;
  background: #d4a017;
  border-radius: 2px;
  margin-bottom: 18px;
}
.mf-desc {
  font-size: .92rem;
  color: #475569;
  line-height: 1.8;
  margin: 0;
  white-space: pre-line;
}

/* ── REGULAR MODAL (vertikal) ── */
.modal-regular {
  max-width: 480px;
  width: 100%;
}
.mr-photo {
  width: 100%;
  height: 240px;
  background: #e8f0e4;
  position: relative;
  overflow: hidden;
}
.mr-photo img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  background: #e8f0e4;
}
.modal-photo-placeholder {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 3rem;
  color: #bbb;
}
.mr-info {
  padding: 28px 24px;
}
.mr-name {
  font-size: 1.3rem;
  font-weight: 800;
  color: #0f172a;
  margin: 0 0 4px;
}
.mr-pos {
  font-size: .88rem;
  font-weight: 600;
  color: #2563eb;
  margin: 0 0 14px;
}
.mr-desc {
  font-size: .88rem;
  color: #475569;
  line-height: 1.7;
  margin: 0;
  white-space: pre-line;
}

@media (max-width: 700px) {
  .modal-featured { flex-direction: column; height: auto; }
  .mf-photo { width: 100%; height: 260px; }
  .mf-info { padding: 28px 24px; }
}

/* ── EMPTY ── */
.empty-state {
  text-align: center; padding: 80px 20px; color: #999;
}
.empty-icon { font-size: 4rem; margin-bottom: 12px; }
.empty-text { font-size: 1rem; }

/* ── RESPONSIVE ── */
@media (max-width: 960px) {
  .hero-content { padding: 48px 40px 60px; }
  .line { font-size: 2.1rem; }
  .team-section { padding: 56px 40px 72px; }
  .featured-card { flex-direction: column; align-items: center; text-align: center; padding: 24px; gap: 20px; min-height: auto; }
  .featured-photo { width: 100%; min-height: 220px; max-width: 280px; border-radius: 12px; }
  .featured-photo img { position: relative; }
  .featured-info { padding: 0; }
  .featured-btn { align-self: center; }
  .team-grid { grid-template-columns: repeat(2, 1fr); gap: 32px 24px; }
}
@media (max-width: 600px) {
  .line { font-size: 1.6rem; }
  .hero-content { padding: 36px 24px 52px; }
  .team-section { padding: 40px 20px 56px; }
  .section-title { font-size: 1.4rem; }
  .team-grid { grid-template-columns: 1fr; gap: 28px; }
}

/* ============ STRUKTUR ORGANISASI (gambar bagan) ============ */
.orgchart-section {
  padding: 0 40px 88px;
  background: #fff;
}
.orgchart-img-wrap {
  max-width: 1200px;
  margin: 0 auto;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 20px rgba(0,0,0,.08);
}
.orgchart-img {
  width: 100%;
  height: auto;
  display: block;
}

@media (max-width: 960px) {
  .orgchart-section { padding: 0 24px 72px; }
}
@media (max-width: 600px) {
  .orgchart-section { padding: 0 16px 56px; }
}

/* ============ FOOTER ============ */
.site-footer {
  width: 100%;
  background-color: #5F9E42;
  padding: 25px 0;
  margin-top: 40px;
}
.footer-container {
  display: flex;
  align-items: center;
  justify-content: center;
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}
.footer-left-content { display: none; }
.footer-copyright { text-align: center; width: 100%; }
.footer-copyright p {
  font-family: Arial, sans-serif;
  font-size: 13px;
  color: #ffffff;
  margin: 0;
  font-weight: normal;
  letter-spacing: 0.3px;
}
</style>