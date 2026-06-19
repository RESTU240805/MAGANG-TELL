<template>
  <section class="relative h-[420px] flex items-end bg-cover bg-center overflow-hidden"
    style="background-image: url('/images/lokasi pabrik.jpeg')">
    <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>
    <div class="absolute top-10 left-10 flex items-center gap-2 opacity-0 anim-fade-right" style="animation-delay: 0.1s">
      <span class="w-1 h-5 bg-green-400 rounded-full inline-block"></span>
      <span class="text-green-400 text-xs font-bold tracking-[3px] uppercase">Sustainability</span>
    </div>
    <div class="relative z-10 max-w-5xl mx-auto px-10 pb-14 w-full">
      <h1 class="text-5xl font-black text-white leading-tight mb-4 opacity-0 anim-fade-up" style="animation-delay: 0.2s">
        Pulp <span class="text-green-400">Process</span>
      </h1>
      <p class="text-gray-300 text-sm max-w-xl leading-relaxed opacity-0 anim-fade-up" style="animation-delay: 0.3s">
        Explore the complete industrial pulp production process from raw logs to finished bales.
      </p>
      <div class="flex items-center gap-2 mt-5 text-xs text-gray-400 opacity-0 anim-fade-up" style="animation-delay: 0.4s">
        <RouterLink to="/" class="hover:text-white cursor-pointer transition">Home</RouterLink>
        <span class="text-gray-600">/</span>
        <span class="hover:text-white cursor-pointer transition">Sustainability</span>
        <span class="text-gray-600">/</span>
        <span class="text-green-400 font-semibold">Pulp Process</span>
      </div>
    </div>
  </section>

  <!-- 3D Interactive Map -->
  <section class="bg-[#0a1208] py-10">
    <div class="max-w-7xl mx-auto px-4">
      <div class="anim-item rounded-xl overflow-hidden shadow-2xl border border-white/10">
        <PulpProcess3D />
      </div>
    </div>
  </section>

  <!-- Process Description -->
  <section class="py-16 bg-white">
    <div class="max-w-4xl mx-auto px-10">
      <h2 class="text-3xl font-black text-green-600 mb-6 anim-item">Pulp Process</h2>
      <p class="text-gray-600 text-sm leading-relaxed mb-6 anim-item">
        There are two main processes in the Pulp and Paper Mill, there are:
      </p>

      <div class="space-y-6">
        <div class="bg-green-50 rounded-xl p-6 border border-green-100 anim-item">
          <h3 class="text-lg font-bold text-gray-900 mb-2">1. Pulp Making Process</h3>
          <p class="text-gray-600 text-sm leading-relaxed">
            Pulp making process is represents the process by which wood or other lignocellulosic material
            that reduced to be fibrous mass. Pulp is produced at PT. TELPP is 450,000 Adt (BKP) / year
            or 1,430 Adt/BKPd.
          </p>
        </div>

        <div class="bg-red-50 rounded-xl p-6 border border-red-100 anim-item">
          <h3 class="text-lg font-bold text-gray-900 mb-2">2. Chemical Recovery Process</h3>
          <p class="text-gray-600 text-sm leading-relaxed">
            Chemical recovery process is the process to recover again cooking chemical (White Liquor)
            that used to cooked wood chip become pulp. White Liquor contains active chemical NaOH and
            Na&#8322;S. Design capacity for Chemical recovery process is 2400 TDS/d solid burning in Recovery Boiler.
          </p>
        </div>
      </div>
    </div>
  </section>

  <!-- Pulp Process Flow Diagram -->
  <section class="py-12 bg-gray-50">
    <div class="max-w-4xl mx-auto px-10">
      <h2 class="text-2xl font-black text-gray-900 text-center mb-8 anim-item">
        Pulp Process <span class="text-green-600">Flow</span>
      </h2>
      <div class="anim-item bg-white rounded-2xl overflow-hidden shadow-sm border border-gray-100">
        <img src="/pulp proces/all_process-2.jpg" alt="Pulp Process Flow Diagram"
          class="w-full h-auto" />
      </div>
    </div>
  </section>

  <!-- Pulp Process Accordion -->
  <section class="py-16 bg-white">
    <div class="max-w-4xl mx-auto px-10">
      <h2 class="text-3xl font-black text-green-600 mb-8 anim-item">Pulp Process</h2>

      <div class="space-y-3">
        <div v-for="(section, i) in pulpSections" :key="i"
          class="border border-gray-200 rounded-xl overflow-hidden anim-item"
          :style="{ animationDelay: `${i * 0.08}s` }">
          <button @click="togglePulpSection(i)"
            class="w-full flex items-center justify-between px-6 py-4 text-left transition-colors cursor-pointer"
            :class="section.open ? 'bg-green-600 text-white' : 'bg-green-600 text-white hover:bg-green-700'">
            <span class="font-bold text-sm">{{ section.title }}</span>
            <svg class="w-5 h-5 transition-transform duration-300" :class="section.open ? 'rotate-180' : ''"
              fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>
          <div v-show="section.open" class="px-6 py-5 bg-gray-50">
            <img v-if="section.image" :src="section.image" :alt="section.title"
              class="w-full h-72 object-cover rounded-lg mb-4" />
            <p class="text-gray-600 text-sm leading-relaxed">{{ section.desc }}</p>
            <div v-if="section.details" class="mt-4 space-y-2">
              <div v-for="(detail, j) in section.details" :key="j"
                class="flex items-start gap-3 text-sm">
                <span class="w-1.5 h-1.5 rounded-full bg-green-500 mt-1.5 shrink-0"></span>
                <span class="text-gray-600">{{ detail }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- Warehouse Photo -->
  <section class="py-12 bg-gray-50">
    <div class="anim-item overflow-hidden shadow-lg">
      <img src="/pulp proces/IMG_6725.jpg" alt="Finished Pulp Bales Warehouse"
        class="w-full h-[500px] object-cover" />
    </div>
  </section>

  <!-- Chemical Recovery Process -->
  <section class="py-16 bg-white">
    <div class="max-w-4xl mx-auto px-10">
      <h2 class="text-3xl font-black text-green-600 mb-4 anim-item">Chemical Recovery Process</h2>
      <p class="text-gray-600 text-sm leading-relaxed mb-8 anim-item">
        The chemical recovery cycle regenerates cooking chemicals and generates energy.
        Black liquor from the washing process is concentrated and burned to recover
        inorganic chemicals while producing steam for power generation.
      </p>

      <div class="flex flex-wrap gap-2 mb-8 anim-item">
        <button v-for="(tab, i) in recoveryTabs" :key="i"
          @click="activeRecoveryTab = i"
          class="px-4 py-2 rounded-full text-xs font-bold tracking-wide transition-all cursor-pointer border"
          :class="activeRecoveryTab === i
            ? 'bg-green-600 text-white border-green-600'
            : 'bg-white text-gray-600 border-gray-200 hover:border-green-400 hover:text-green-600'">
          {{ tab }}
        </button>
      </div>

      <div class="bg-gray-50 rounded-xl overflow-hidden border border-gray-100 anim-item">
        <img v-if="recoveryData[activeRecoveryTab].image"
          :src="recoveryData[activeRecoveryTab].image"
          :alt="recoveryData[activeRecoveryTab].title"
          class="w-full h-72 object-cover" />
        <div class="p-6">
          <h3 class="text-lg font-bold text-gray-900 mb-3">{{ recoveryData[activeRecoveryTab].title }}</h3>
          <p class="text-gray-600 text-sm leading-relaxed mb-4">{{ recoveryData[activeRecoveryTab].desc }}</p>
          <div class="grid grid-cols-2 gap-4">
            <div v-for="(spec, j) in recoveryData[activeRecoveryTab].specs" :key="j"
              class="bg-white rounded-lg p-4 border border-gray-100">
              <div class="text-xs text-gray-400 mb-1">{{ spec.label }}</div>
              <div class="text-sm font-bold text-gray-800">{{ spec.value }}</div>
            </div>
          </div>
        </div>
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
import { ref, reactive, onMounted, onUnmounted, nextTick } from 'vue'
import { RouterLink } from 'vue-router'
import PulpProcess3D from '@/components/PulpProcess3D.vue'

const pulpSections = reactive([
  {
    title: 'Wood and Chip Handling',
    open: true,
    image: '/pulp proces/Chipping-Proses.jpg',
    desc: 'Raw logs from sustainable industrial plantations (Acacia Mangium & Eucalyptus Pellita) are transported to the mill, stored in wet ponds, then processed through debarking and chipping.',
    details: [
      'Log Storage Ponds — logs kept wet to preserve fiber quality',
      'Debarking Drum — rotary drum strips bark (bark used as biomass fuel)',
      'Disk Chipper — reduces debarked logs into uniform 25mm chips',
      'Chip Screening — separates oversized chips (rechipped) and fines (fuel)'
    ]
  },
  {
    title: 'Fiber Line',
    open: false,
    image: '/pulp proces/FiberLine.jpg',
    desc: 'Wood chips are cooked under high pressure with white liquor to separate cellulose fibers from lignin, then washed and delignified.',
    details: [
      'Chip Bin & Digester — Kamyr continuous digester at 170\u00b0C with NaOH + Na\u2082S',
      'Blow Tank — receives cooked pulp at atmospheric pressure',
      'Brownstock Washing — multi-stage washers remove black liquor',
      'O\u2082 Delignification — reduces kappa number 40-50% using oxygen'
    ]
  },
  {
    title: 'Pulp Machine',
    open: false,
    image: '/pulp proces/pulp-machine.jpg',
    desc: 'The bleached pulp slurry is formed into sheets on a Fourdrinier wire, pressed, and steam-dried into market pulp.',
    details: [
      'Forming Section — diluted pulp (0.5-1% consistency) deposited on wire',
      'Press Section — water removed by pressing between felt rolls',
      'Dryer Section — steam-heated cylinders dry pulp to ~90% solids',
      'Production capacity: 450,000 ADt/year (1,430 Adt/day)'
    ]
  },
  {
    title: 'Warehouse',
    open: false,
    image: '/pulp proces/IMG_6725.jpg',
    desc: 'Finished pulp sheets are cut, stacked, and compressed into 250 kg bales wrapped in kraft liner for storage and export.',
    details: [
      'Bale Press — sheets compressed into 250 kg bales',
      'Wrapping — bales wrapped in kraft liner for protection',
      'Climate-controlled storage — temperature and humidity regulated',
      'Conveyor belt to jetty — bales loaded onto cargo ships for export'
    ]
  }
])

function togglePulpSection(index) {
  pulpSections[index].open = !pulpSections[index].open
}

const activeRecoveryTab = ref(0)
const recoveryTabs = ['Evaporator', 'Recovery Boiler', 'Recausticizing', 'Lime Kiln']

const recoveryData = [
  {
    title: 'Evaporator Plant (7-Effect)',
    image: '/pulp proces/Epavorator.jpg',
    desc: 'Multiple-effect evaporator concentrates weak black liquor from 15% to 72% dissolved solids for firing in the Recovery Boiler. The 7-effect system maximizes steam economy.',
    specs: [
      { label: 'Inlet Concentration', value: '15% Dissolved Solids' },
      { label: 'Outlet Concentration', value: '72% Dissolved Solids' },
      { label: 'Effect Stages', value: '7-Effect System' },
      { label: 'Heat Source', value: 'Exhaust Steam from Boilers' }
    ]
  },
  {
    title: 'Recovery Boiler',
    image: '/pulp proces/Recovery-Boiler.jpg',
    desc: 'Burns concentrated black liquor (70% DS) to recover inorganic smelt and generate high-pressure steam. The smelt dissolves in green liquor for recausticizing.',
    specs: [
      { label: 'Design Capacity', value: '2,400 TDS/day' },
      { label: 'Fuel', value: 'Concentrated Black Liquor' },
      { label: 'Products', value: 'Smelt + High-Pressure Steam' },
      { label: 'Steam Usage', value: 'Power Generation + Process' }
    ]
  },
  {
    title: 'Recausticizing Plant',
    image: '/pulp proces/Recautisizing.jpg',
    desc: 'Green liquor (Na\u2082CO\u2083 + Na\u2082S) is reacted with Ca(OH)\u2082 to regenerate White Liquor (NaOH + Na\u2082S) for the cooking cycle. Lime mud (CaCO\u2083) is sent to the lime kiln.',
    specs: [
      { label: 'Input', value: 'Green Liquor (Na\u2082CO\u2083 + Na\u2082S)' },
      { label: 'Reagent', value: 'Ca(OH)\u2082 (Slaked Lime)' },
      { label: 'Output', value: 'White Liquor (NaOH + Na\u2082S)' },
      { label: 'By-product', value: 'Lime Mud (CaCO\u2083) \u2192 Kiln' }
    ]
  },
  {
    title: 'Lime Kiln',
    image: '/pulp proces/Lime-Kiln.jpg',
    desc: 'Rotary kiln (~150m) reburns CaCO\u2083 lime mud at 1000\u00b0C to regenerate CaO quicklime for the recausticizing cycle. Completes the chemical recovery loop.',
    specs: [
      { label: 'Length', value: '~150 meters' },
      { label: 'Temperature', value: '1,000\u00b0C' },
      { label: 'Input', value: 'Lime Mud (CaCO\u2083)' },
      { label: 'Output', value: 'Quicklime (CaO) \u2192 Recausticizing' }
    ]
  }
]

let observer = null

onMounted(async () => {
  await nextTick()
  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('anim-visible')
        observer.unobserve(entry.target)
      }
    })
  }, { threshold: 0.15 })
  document.querySelectorAll('.anim-item').forEach(el => observer.observe(el))
})

onUnmounted(() => {
  if (observer) {observer.disconnect()}
})
</script>

<style scoped>
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
.anim-item {
  opacity: 0;
  transform: translateY(30px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-item.anim-visible {
  opacity: 1;
  transform: translateY(0);
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