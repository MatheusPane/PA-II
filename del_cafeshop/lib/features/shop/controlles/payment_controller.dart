import 'dart:convert';
import 'package:http/http.dart' as http;

Future<String?> createTransaction({
  required String orderId,
  required int amount,
  required String customerName,
}) async {
  final url = Uri.parse('http://172.27.69.192:8000/admin/payment'); // ganti dengan IP server kamu

  final response = await http.post(
    url,
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({
      'orderId': orderId,
      'amount': amount,
      'customerName': customerName,
    }),
  );

  if (response.statusCode == 200) {
    final data = jsonDecode(response.body);
    return data['snapToken'];
  } else {
    print('Gagal membuat transaksi: ${response.body}');
    return null;
  }
}
