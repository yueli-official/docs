import { createUiPreset } from '@yueli/ui/theme'

const preset = createUiPreset({ primary: 'indigo' })
const fieldBorder = 'docs-field-border'

export default defineAppConfig({
  ui: {
    ...preset.ui,
    input: {
      slots: { base: fieldBorder },
    },
    inputNumber: {
      slots: { base: fieldBorder },
    },
    textarea: {
      slots: { base: fieldBorder },
    },
    select: {
      slots: { base: fieldBorder },
    },
    selectMenu: {
      slots: { base: fieldBorder },
    },
  },
})
