import 'package:flutter/material.dart';
import 'package:webview_flutter/webview_flutter.dart';

class PaymentPage extends StatefulWidget {
  final String snapToken;

  const PaymentPage({super.key, required this.snapToken});

  @override
  State<PaymentPage> createState() => _PaymentPageState();
}

class _PaymentPageState extends State<PaymentPage> {
  late final WebViewController _controller;

  @override
  void initState() {
    super.initState();

    print('Snap Token: ${widget.snapToken}'); // Tampilkan snapToken ke console
    final snapUrl = 'https://app.midtrans.com/snap/v2/vtweb/${widget.snapToken}';

    _controller = WebViewController()
      ..setJavaScriptMode(JavaScriptMode.unrestricted)
      ..loadRequest(Uri.parse(snapUrl));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Bayar Sekarang')),
      body: WebViewWidget(controller: _controller),
    );
  }
}
