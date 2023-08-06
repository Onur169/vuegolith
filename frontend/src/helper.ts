export function formatBytes(bytes: number): string {
  if (bytes >= 1000000000) {
    return (bytes / 1000000000).toFixed(2) + 'GB';
  } else if (bytes >= 1000000) {
    return (bytes / 1000000).toFixed(2) + 'MB';
  } else if (bytes >= 1000) {
    return (bytes / 1000).toFixed(2) + 'KB';
  } else {
    return bytes + 'B';
  }
}

export function initArrayWithBooleans(length: number) {
  return Array(length).fill(false);
}

export function updateHoveringState(array: boolean[], index: number, hovering: boolean): boolean[] {
  const newArray = [...array];
  newArray[index] = hovering;
  return newArray;
}
