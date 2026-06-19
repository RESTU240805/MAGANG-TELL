<template>

  <!-- ── HERO ── -->
  <section class="relative h-[420px] flex items-end bg-cover bg-center overflow-hidden"
    style="background-image: url('/images/people3.jpeg')">
    <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>

    <div class="absolute top-10 left-10 flex items-center gap-2 opacity-0 anim-fade-right" style="animation-delay: 0.1s">
      <span class="w-1 h-5 bg-green-400 rounded-full"></span>
      <span class="text-green-400 text-xs font-bold tracking-[3px] uppercase">Sustainability</span>
    </div>

    <div class="relative z-10 max-w-5xl mx-auto px-10 pb-14 w-full">
      <h1 class="text-5xl font-black text-white leading-tight mb-4 opacity-0 anim-fade-up" style="animation-delay: 0.2s">
        People <span class="text-green-400">Development</span>
      </h1>
      <p class="text-gray-300 text-sm max-w-xl leading-relaxed opacity-0 anim-fade-up" style="animation-delay: 0.3s">
        Preparing the best talent to achieve the highest performance and business sustainability.
      </p>
      <div class="flex items-center gap-2 mt-5 text-xs text-gray-400 opacity-0 anim-fade-up" style="animation-delay: 0.4s">
        <RouterLink to="/" class="hover:text-white cursor-pointer transition">Home</RouterLink>
        <span class="text-gray-600">/</span>
        <span class="hover:text-white cursor-pointer transition">Sustainability</span>
        <span class="text-gray-600">/</span>
        <span class="text-green-400 font-semibold">People Development</span>
      </div>
    </div>
  </section>

  <!-- ── PHOTO GRID ── -->
  <section class="py-16 bg-white">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-3 gap-3 h-64 rounded-2xl overflow-hidden">
        <div class="col-span-1 overflow-hidden anim-item">
          <img src="/images/Training-First-Aid.jpg"
            alt="Training" class="w-full h-full object-cover hover:scale-105 transition-transform duration-500"/>
        </div>
        <div class="col-span-1 overflow-hidden anim-item" style="animation-delay: 0.1s">
          <img src="/images/Management-Safety-Patrol.jpg"
            alt="Workshop" class="w-full h-full object-cover hover:scale-105 transition-transform duration-500"/>
        </div>
        <div class="col-span-1 overflow-hidden anim-item" style="animation-delay: 0.2s">
          <img src="/images/Industrial-Hygiene.jpg"
            alt="Team" class="w-full h-full object-cover hover:scale-105 transition-transform duration-500"/>
        </div>
      </div>
    </div>
  </section>

  <!-- ── INTRO TEXT ── -->
  <section class="pb-16 bg-white">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-5 gap-12">

        <div class="col-span-3 space-y-4 text-gray-500 text-sm leading-relaxed anim-slide-right">
          <p>
            People development is carried out in order to prepare the best talent to be able to show the
            best performance and to achieve the business targets they carry. Talking about business is not
            only about running but also how they manage the business even more and sustain it in the future.
          </p>
          <p>
            PT Tanjungenim Lestari Pulp and Paper are committed to developing competent, motivated and
            integrated information systems to achieve a competitive organization. The commitment is realized
            in several activities and achievements as follows:
          </p>
        </div>

        <div class="col-span-2 bg-green-50 border border-green-100 rounded-2xl p-6 flex flex-col justify-center anim-slide-left">
          <div class="w-10 h-10 rounded-xl bg-green-100 flex items-center justify-center mb-4">
            <span class="text-green-600 text-lg">🌱</span>
          </div>
          <p class="text-green-800 font-bold text-base leading-snug mb-2">
            Competent. Motivated. Integrated.
          </p>
          <p class="text-green-600 text-xs leading-relaxed">
            Building people who are ready to sustain and grow the business for the long term.
          </p>
        </div>

      </div>
    </div>
  </section>

  <!-- ── COMMITMENT LIST ── -->
  <section class="py-16 bg-[#14532d]">
    <div class="max-w-5xl mx-auto px-10">

      <div class="text-center mb-12">
        <span class="inline-block bg-green-800 text-green-200 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-3 anim-item">
          Our Commitments
        </span>
        <h2 class="text-3xl font-black text-white anim-item">Development Pillars</h2>
        <p class="text-green-200/60 text-sm mt-2 max-w-md mx-auto anim-item">
          Five core programs that shape TEL's approach to people development.
        </p>
      </div>

      <div class="space-y-4">
        <div v-for="(item, index) in commitments" :key="index"
          class="flex gap-5 bg-white/10 backdrop-blur-sm border border-white/10 rounded-2xl p-6 hover:bg-white/15 transition group anim-item"
          :style="{ animationDelay: (index * 0.1) + 's' }">

          <div class="w-10 h-10 rounded-xl bg-green-500 text-white font-black text-sm flex items-center justify-center flex-shrink-0 group-hover:bg-green-400 transition">
            {{ String(index + 1).padStart(2, '0') }}
          </div>

          <div>
            <h3 class="text-white font-bold text-sm mb-1">{{ item.title }}</h3>
            <p class="text-green-200/60 text-sm leading-relaxed">{{ item.desc }}</p>
          </div>
        </div>
      </div>

    </div>
  </section>

  <!-- ── IMAGE SLIDER ── -->
  <section class="py-20 bg-white">
    <div class="max-w-5xl mx-auto px-10">

      <div class="text-center mb-10">
        <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-3 anim-item">
          Gallery
        </span>
        <h2 class="text-3xl font-black text-gray-900 anim-item">People in Action</h2>
      </div>

      <!-- Slider wrapper -->
      <div class="relative rounded-2xl overflow-hidden shadow-xl aspect-video bg-gray-100 anim-item">

        <img
          v-for="(slide, i) in slides"
          :key="i"
          :src="slide.url"
          :alt="slide.caption"
          class="absolute inset-0 w-full h-full object-cover"
          :style="{ opacity: activeSlide === i ? 1 : 0, transition: 'opacity 0.8s ease-in-out' }"
        />

        <!-- Caption -->
        <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6">
          <p class="text-white text-sm font-semibold">{{ slides[activeSlide].caption }}</p>
          <p class="text-gray-300 text-xs mt-1">{{ activeSlide + 1 }} / {{ slides.length }}</p>
        </div>

        <!-- Prev -->
        <button @click="manualNav(prev)"
          class="absolute left-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-black/50 text-white flex items-center justify-center hover:bg-black/70 transition backdrop-blur-sm">
          ‹
        </button>

        <!-- Next -->
        <button @click="manualNav(next)"
          class="absolute right-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-black/50 text-white flex items-center justify-center hover:bg-black/70 transition backdrop-blur-sm">
          ›
        </button>
      </div>

      <!-- Dot Pagination -->
      <div class="flex justify-center gap-2 mt-6">
        <button
          v-for="(_, i) in slides" :key="i"
          @click="manualNav(() => activeSlide = i)"
          class="transition-all duration-300 rounded-full cursor-pointer"
          :class="activeSlide === i ? 'bg-green-600 w-6 h-2.5' : 'bg-gray-200 hover:bg-gray-300 w-2.5 h-2.5'">
        </button>
      </div>
    </div>
  </section>

  <section class="contact-section">
      <div class="contact-container">
        <div class="contact-header anim-item">
          <span class="contact-label">GET IN TOUCH</span>
          <h2 class="contact-title">Our Offices</h2>
        </div>

        <div class="contact-grid">
          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/jakarta.jpeg" alt="Jakarta Office" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Jakarta</h3>
              <p class="contact-address">Menara Astra 22nd floor – Zona D, Jalan Jenderal Sudirman Kav. 5-6 Kel. Karet Tengsin, Kec. Tanah Abang</p>
              <a class="contact-phone" href="tel:+622186656809">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                +62 21 8665 6809 / 8665 6810
              </a>
            </div>
          </div>

          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/lokasi pabrik.jpeg" alt="Mill Site" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Mill Site</h3>
              <p class="contact-address">Desa Banuayu, Kec. Empat Petulai Dangku, Kab. Muara Enim, Sumatera Selatan</p>
              <a class="contact-phone" href="tel:+62713324150">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                (62) (713) 324150 – 324160
              </a>
            </div>
          </div>

          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/palembang.jpeg" alt="Palembang Office" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Palembang</h3>
              <p class="contact-address">Ruko Blok I/29, Komplek PTC Mall. Jl. R. Soekamto Palembang 30114, Sumatera Selatan</p>
              <a class="contact-phone" href="tel:+62711382409">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                (62) (711) 382409
              </a>
            </div>
          </div>

          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/tarahan.jpeg" alt="Tarahan Port" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Tarahan</h3>
              <p class="contact-address">Jl. Soekarno Hatta Km. 14, Batu Serampok Kel. Srengsem Kec. Panjang, Bandar Lampung</p>
              <a class="contact-phone" href="tel:+62721342311">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                (62) (721) 34231, 31318
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="footer-info-section">
      <div class="footer-info-container">
        <div class="footer-info-grid">

          <div class="footer-info-col anim-item">
            <p class="footer-about-text">
              PT. Tanjungenim Lestari Pulp and Paper (PT. TEL), is one of the most exciting pulp mills in Indonesia today and the only pulp mill in the world to produce high-quality, bleached-hardwood Kraft pulp with 100 percent plantation grown Acacia mangium and Eucalyptus Pellita trees.
            </p>
          </div>

          <div class="footer-info-col footer-logos-col anim-item">
            <img src="/images/logosatu.jpeg" alt="Yayasan Pendidikan Tanjungenim Lestari - SMP, SD, PGTK Lematang Lestari" class="footer-logo-combined" />
          </div>

          <div class="footer-info-col footer-contact-col anim-item">
            <div class="footer-contact-row">
              <svg class="footer-contact-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 1 1 18 0z"/><circle cx="12" cy="10" r="3"/></svg>
              <span>Kab. Muara Enim, Sumatera Selatan.</span>
            </div>
            <div class="footer-contact-row">
              <svg class="footer-contact-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
              <span>(+62) 713-324-150</span>
            </div>
            <div class="footer-contact-row">
              <svg class="footer-contact-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 6v6l4 2"/></svg>
              <span>Mon-Fri: 8:00 – 17:00</span>
            </div>
          </div>

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

</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { RouterLink } from 'vue-router'

const activeSlide = ref(0)
let sliderTimer = null

const commitments = [
  {
    title: 'Skills-Based Leadership Coaching',
    desc: 'Achieve a high performing organization by developing skills-based Coaching for leadership styles through a comprehensive Supervisory / Management / Leadership Program.'
  },
  {
    title: 'Succession Planning & Career Path',
    desc: 'Develop and implement succession plans for critical positions to ensure the regeneration of employees through the Career Path and Talent Management.'
  },
  {
    title: 'Community & Scholarship Programs',
    desc: 'Prepare a generation of quality sources from internal employees as well as from the village community through scholarships, internship, community education, school foundation programs and infrastructure improvement.'
  },
  {
    title: 'Technology & Information Skills',
    desc: 'Improve knowledge and skills related to technology and information across all levels of the organization.'
  },
  {
    title: 'Rules, Competencies & Knowledge',
    desc: 'Meet the rules and competencies with good knowledge and skills to maintain professional standards throughout the company.'
  },
]

const slides = [
  { url: '/images/Training-First-Aid.jpg', caption: 'First Aid Training Session' },
  { url: '/images/Management-Safety-Patrol.jpg', caption: 'Safety Patrol Management' },
  { url: '/images/Industrial-Hygiene.jpg', caption: 'Team Safety Workshop' },
  { url: '/images/people1.jpeg', caption: 'Industrial Hygiene Program' },
  { url: '/images/people2.jpeg', caption: 'People Development Initiative' },
]

const next = () => { activeSlide.value = (activeSlide.value + 1) % slides.length }
const prev = () => { activeSlide.value = (activeSlide.value - 1 + slides.length) % slides.length }
const startSlider = () => { sliderTimer = setInterval(next, 5000) }
const stopSlider = () => { clearInterval(sliderTimer) }
const manualNav = (fn) => { stopSlider(); fn(); startSlider() }

// ─── Scroll Observer ──────────────────────────────────────────
let observer = null

onMounted(async () => {
  await nextTick()
  initObserver()
  startSlider()
})

onUnmounted(() => {
  if (observer) {observer.disconnect()}
  stopSlider()
})

const initObserver = () => {
  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('anim-visible')
      } else {
        entry.target.classList.remove('anim-visible')
      }
    })
  }, { 
    threshold: 0.15,
    rootMargin: '0px 0px -40px 0px'
  })

  document.querySelectorAll('.anim-item, .anim-slide-right, .anim-slide-left').forEach(el => {
    observer.observe(el)
  })
}
</script>

<style scoped>
/* Hero animations */
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(24px); }
  to { opacity: 1; transform: translateY(0); }
}
.anim-fade-up {
  animation: fadeUp 0.8s cubic-bezier(0.22, 1, 0.36, 1) forwards;
}

@keyframes fadeRight {
  from { opacity: 0; transform: translateX(-24px); }
  to { opacity: 1; transform: translateX(0); }
}
.anim-fade-right {
  animation: fadeRight 0.8s cubic-bezier(0.22, 1, 0.36, 1) forwards;
}

/* Scroll animations */
.anim-item {
  opacity: 0;
  transform: translateY(30px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-item.anim-visible {
  opacity: 1;
  transform: translateY(0);
}

.anim-slide-right {
  opacity: 0;
  transform: translateX(-40px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-slide-right.anim-visible {
  opacity: 1;
  transform: translateX(0);
}

.anim-slide-left {
  opacity: 0;
  transform: translateX(40px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-slide-left.anim-visible {
  opacity: 1;
  transform: translateX(0);
}


.contact-container {
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}

.contact-header {
  text-align: center;
  margin-bottom: 48px;
}

.contact-label {
  color: #5F9E42;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.25em;
  text-transform: uppercase;
  display: block;
  margin-bottom: 10px;
}

.contact-title {
  font-size: 32px;
  font-weight: 800;
  color: #18301f;
  letter-spacing: -0.02em;
}

.contact-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

@media (max-width: 1024px) {
  .contact-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .contact-grid {
    grid-template-columns: 1fr;
  }
}

.contact-card {
  background: #ffffff;
  border: 1px solid #dcecd2;
  border-radius: 20px;
  padding: 28px 24px;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.contact-card:hover {
  transform: translateY(-6px);
  border-color: #a9d68c;
  box-shadow: 0 16px 32px rgba(95, 158, 66, 0.12);
}

.contact-icon {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  background: #eaf6e1;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 18px;
  color: #5F9E42;
}

.contact-icon svg {
  width: 22px;
  height: 22px;
}

.contact-city {
  font-size: 17px;
  font-weight: 700;
  color: #18301f;
  margin-bottom: 10px;
}

.contact-address {
  font-size: 13px;
  color: #6b7b66;
  line-height: 1.6;
  margin-bottom: 16px;
  font-weight: 400;
}

.contact-phone {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13.5px;
  font-weight: 600;
  color: #3b6d11;
  text-decoration: none;
  transition: color 0.3s ease;
}

.contact-phone:hover {
  color: #27500a;
}

.contact-phone svg {
  width: 15px;
  height: 15px;
  flex-shrink: 0;
}

.contact-section {
  background: linear-gradient(180deg, #ffffff 0%, #eef7e9 100%);
  padding: 80px 0 70px;
}

.contact-container {
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}

.contact-header {
  text-align: center;
  margin-bottom: 48px;
}

.contact-label {
  color: #5F9E42;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.25em;
  text-transform: uppercase;
  display: block;
  margin-bottom: 10px;
}

.contact-title {
  font-size: 32px;
  font-weight: 800;
  color: #18301f;
  letter-spacing: -0.02em;
}

.contact-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

@media (max-width: 1024px) {
  .contact-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .contact-grid {
    grid-template-columns: 1fr;
  }
}

.contact-card {
  background: #ffffff;
  border: 1px solid #dcecd2;
  border-radius: 20px;
  overflow: hidden;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.contact-card:hover {
  transform: translateY(-6px);
  border-color: #a9d68c;
  box-shadow: 0 16px 32px rgba(95, 158, 66, 0.15);
}

.contact-image {
  width: 100%;
  aspect-ratio: 4 / 3;
  overflow: hidden;
  background: #eaf6e1;
}

.contact-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  transition: transform 0.6s cubic-bezier(0.16, 1, 0.3, 1);
}

.contact-card:hover .contact-image img {
  transform: scale(1.06);
}

.contact-body {
  padding: 20px 18px 22px;
}

.contact-city {
  font-size: 18px;
  font-weight: 800;
  color: #1a4d8f;
  margin-bottom: 12px;
  letter-spacing: -0.01em;
}

.contact-address {
  font-size: 13px;
  color: #5a6b56;
  line-height: 1.6;
  margin-bottom: 16px;
  font-weight: 400;
}

.contact-phone {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13.5px;
  font-weight: 500;
  color: #2c3a2a;
  text-decoration: none;
  transition: color 0.3s ease;
}

.contact-phone:hover {
  color: #3b6d11;
}

.contact-phone svg {
  width: 15px;
  height: 15px;
  flex-shrink: 0;
  color: #5F9E42;
}
/* Responsif untuk Layar Handphone (Mobile) */
@media (max-width: 768px) {
  .footer-container {
    flex-direction: column;
    gap: 15px;
    padding: 0 20px;
  }
  
  .footer-copyright {
    text-align: center; /* Di HP otomatis pindah ke tengah bawah agar rapi */
    width: 100%;
  }
  
  .footer-copyright p {
    font-size: 12px; /* Ukuran sedikit disesuaikan untuk layar kecil */
  }
}

.footer-info-section {
  background: #ffffff;
  border-top: 1px solid #e4efdb;
  padding: 56px 0;
}

.footer-info-container {
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}
.footer-info-grid {
  display: grid;
  grid-template-columns: 1fr 1.1fr 0.9fr;
  gap: 48px;
  align-items: start;
}

.footer-info-col {
  display: flex;
  flex-direction: column;
}

.footer-about-text {
  margin: 0;
  font-size: 14px;
  color: #7a8a76;
  line-height: 1.75;
  font-weight: 400;
}

.footer-logos-col {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0;
  margin: 0;
  padding: 0;
}

.footer-logos-col {
  display: flex;
  align-items: flex-start;
  justify-content: center;
  margin: 0;
  padding: 0;
}

.footer-logo-combined {
  width: 100%;
  max-width: 380px;
  height: auto;
  object-fit: contain;
  display: block;
}

.footer-contact-col {
  display: flex;
  flex-direction: column;
  gap: 20px;
  max-width: 280px;
  margin: 0;
  padding: 0;
}

.footer-contact-row {
  display: flex;
  align-items: center;
  gap: 14px;
  font-size: 14px;
  color: #5a6b56;
  font-weight: 400;
  line-height: 1.4;
  margin: 0;
}
.footer-contact-icon {
  width: 19px;
  height: 19px;
  flex-shrink: 0;
  color: #5F9E42;
}

/* ============ FOOTER ============ */
.site-footer {
  width: 100%;
  /* Warna hijau tua disesuaikan dengan identitas PT TEL pada gambar */
  background-color: #5F9E42; 
  padding: 25px 0; /* Memberikan ruang tinggi baris yang pas */
  margin-top: 40px; /* Jarak pemisah dari konten di atasnya */
}

.footer-container {
  display: flex;
  align-items: center;
  justify-content: center; /* DIUBAH: Membuat konten berkumpul di tengah */
  max-width: 1140px; /* Menyelaraskan lebar dengan konten atas website */
  margin: 0 auto;
  padding: 0 40px; /* Jarak aman agar teks tidak terlalu menempel ke pinggir layar */
}

/* Kolom Kiri Footer */
.footer-left-content {
  display: none; /* DIUBAH: Disembunyikan karena kosong, agar tidak memakan space */
}

/* Kolom Kanan Footer (Teks Copyright) */
.footer-copyright {
  text-align: center; /* DIUBAH: Memastikan teks rata tengah */
  width: 100%; /* Memastikan block mengambil ruang penuh untuk fungsi centering */
}

.footer-copyright p {
  font-family: Arial, sans-serif;
  font-size: 13px; /* Ukuran teks kecil, tipis, dan normal untuk footer */
  color: #ffffff;  /* Warna teks putih bersih agar kontras dengan background hijau */
  margin: 0;
  font-weight: normal;
  letter-spacing: 0.3px;
}
</style>