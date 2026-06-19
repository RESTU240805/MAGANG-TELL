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

        <div v-else-if="members.length" class="team-list">
          <div
            v-for="(m, i) in members"
            :key="m.ID"
            class="team-card"
            :style="{ animationDelay: (i * 0.12) + 's' }"
          >
            <div class="team-photo">
              <div class="photo-placeholder">👤</div>
              <img
                v-if="m.photo_path"
                :src="getImageUrl(m.photo_path)"
                :alt="m.name"
                class="team-img"
              />
            </div>
            <div class="team-info">
              <h3 class="team-name">{{ m.name }}</h3>
              <p class="team-position">{{ m.position }}</p>
              <p v-if="m.description" class="team-desc">{{ m.description }}</p>
            </div>
          </div>
        </div>

        <div v-else class="empty-state">
          <p class="empty-icon">👥</p>
          <p class="empty-text">Belum ada data anggota tim.</p>
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
import { ref, onMounted } from 'vue'
import api from '../services/api'

const BASE_URL = 'http://localhost:8080'
const members = ref([])
const loading = ref(true)

const getImageUrl = (path) => {
  if (!path) return ''
  if (path.startsWith('http')) return path
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

onMounted(async () => {
  try {
    const res = await api.get('/team-members')
    members.value = res.data?.data || []
  } catch {
    members.value = []
  } finally {
    loading.value = false
  }
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
.section-inner { max-width: 1100px; margin: 0 auto; }

.section-title {
  font-size: 1.8rem;
  font-weight: 800;
  color: #1a1a1a;
  margin-bottom: 40px;
}

/* ── HORIZONTAL CARD ── */
.team-card {
  display: flex;
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 16px;
  overflow: hidden;
  margin-bottom: 28px;
  transition: box-shadow .4s, border-color .3s, transform .3s;
  animation: fadeInUp .6s cubic-bezier(.22,1,.36,1) both;
}
.team-card:hover {
  box-shadow: 0 16px 48px rgba(0,0,0,.08);
  border-color: #c8d6c0;
  transform: translateY(-3px);
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(24px); }
  to { opacity: 1; transform: none; }
}

.team-photo {
  width: 380px;
  min-height: 320px;
  flex-shrink: 0;
  overflow: hidden;
  background: #e8f0e4;
  position: relative;
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
  flex: 1;
  padding: 36px 40px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.team-name {
  font-size: 1.35rem;
  font-weight: 800;
  color: #1a1a1a;
  margin-bottom: 4px;
  letter-spacing: -.01em;
}
.team-position {
  font-size: .9rem;
  color: #2563eb;
  font-weight: 500;
  margin-bottom: 20px;
}
.team-desc {
  font-size: .9rem;
  color: #555;
  line-height: 1.8;
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
  .team-card { flex-direction: column; }
  .team-photo { width: 100%; min-height: 280px; max-height: 360px; }
  .team-info { padding: 28px 24px 32px; }
}
@media (max-width: 600px) {
  .line { font-size: 1.6rem; }
  .hero-content { padding: 36px 24px 52px; }
  .team-section { padding: 40px 20px 56px; }
  .section-title { font-size: 1.4rem; }
  .team-photo { min-height: 240px; }
  .team-info { padding: 24px 20px 28px; }
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
