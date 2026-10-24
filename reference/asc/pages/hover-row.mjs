// The overview's plain table's row under the pointer: read with the pointer
// resting on a row.
const ROW = ['background-color', 'border-bottom-color', 'cursor'];
export default {
  path: '/apps/<app>/analytics/overview',
  probes: [
    { name: 'table.row', select: 'table[class*="rt-Table"] tr:hover', props: ROW, box: true },
    { name: 'table.cell', select: 'table[class*="rt-Table"] tr:hover td', props: ['background-color', 'color', 'border-bottom-color'], box: true },
  ],
};
