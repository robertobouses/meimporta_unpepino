<script setup>
import WelcomeItem from './WelcomeItem.vue'
import BusquedasIcon from './icons/IconBusquedas.vue'
import CalculadorIcon from './icons/IconCalculador.vue'
import DefineTuCropIcon from './icons/IconDefineTuCrop.vue'
import Pests       Icon from './icons/IconPests       .vue'
import LecturasIcon from './icons/IconLecturas.vue'
import 'bootstrap-icons/font/bootstrap-icons.css';
</script>

<template>
  <WelcomeItem>
    <template #icon>
      <BusquedasIcon />
    </template>
    <template #heading>Búsquedas</template>

    Encuentra la información sobre tu crop, o filtra y encuentra el crop que ni sospechabas necesitar, nuestros filtros y algoritmos trabajan para ti!
    
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <CalculadorIcon />
    </template>
    <template #heading>Calculador de plantas</template>

    Consigue saber las plantas que necesitas para tu crop, o los metros que necesitas para tus plantines, obtén la distancia de plantación...
   
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <DefineTuCropIcon />
    </template>
    <template #heading>Define tu propio crop</template>

    Un espacio muy tuyo, introduce los datos de tu plantación y tendrás en la palma de tu mano el diario de crop que necesitas. No te olvidarás de regar más, sabrás cuándo plantar, cuánto fertilizar exactamente y tendrás hasta predicciones de cuantos fruits vas a obtener. Todo adaptado a tu crop real.
    
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <Pests       Icon />
    </template>
    <template #heading>Pests        y problemas</template>

    Porque nada es perfecto, los imprevistos y la naturaleza son hermanos, la mejor arma la sabiduría, conoce todo tipo de pests        que pueden afectar a tu crop, aprende también sobre otros problemas meteorológicos, hongos, de nutrición e incluso físicos. Cuidamos de tus plantas.
    
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <LecturasIcon />
    </template>
    <template #heading>Lecturas</template>

    Si eres curioso, aún tenemos más, lee curiosidades, sobre modos de crop, un poco de historia y sobre todo sobre agricultura.
   </WelcomeItem>
</template>
