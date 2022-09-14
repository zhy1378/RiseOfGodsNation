package rogd;

/**
 * @author Hans ZHANG, OoKai Tech.
 * @version 1.0.0
 * @since 2018-01-01 07:51
 */
public class RiseOfGodsNation {

	public static void main(String... args) {
		double e = 2.71828183;

		// (int) (e ** 层 + e * (等级 - 层等级) * (等级 - 层等级))
		for (int a = 0, b = 1, c, layer = 0; layer < 20; layer++) {
			c = a + b;
			a = b;
			b = c;
//			System.out.println("层: " + layer + ", 等级: " + a + ", " + (isPrime(a) ? '质' : ' '));
//			System.out.println("Layer: " + i + ", level: " + a + ", Energy = " + (int) Math.pow(e, i) + ", origin = " + Math.pow(e, i));
			System.out.println("层= " + layer + ", \t级= " + a + (isPrime(a) ? '质' : ' ') + ", 能量 = " + (int) Math.pow(e, layer) + ", 取整自 " + Math.pow(e, layer));
			for (int level = a + 1; level < c; level++) {
				System.out.println("\t\t级= " + level + ", \t能量= " + (int) ((Math.pow(e, layer) + e * (level - a) * (level - a))));
			}
		}
	}

	private static boolean isPrime(int num) {
		if (num < 2) return false;
		if (num == 2) return true;
		if (num % 2 == 0) return false;
		for (int i = 3; i * i <= num; i += 2)
			if (num % i == 0) return false;
		return true;
	}
}
