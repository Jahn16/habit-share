import palette from 'google-palette';
export const getColors = (userPalette: string | undefined, qty: number) => {
	const colorPalette = qty <= 9 ? userPalette || 'cb-Blues' : 'tol';
	const colorQty = qty < 9 ? 9 : qty;

	const colors: string[] = palette(colorPalette, colorQty);
	return colors.reverse();
};
