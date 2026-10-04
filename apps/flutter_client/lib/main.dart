import 'package:flutter/material.dart';

void main() {
  runApp(const CryptoCloudMiningApp());
}

class CryptoCloudMiningApp extends StatelessWidget {
  const CryptoCloudMiningApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Crypto Cloud Mining',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.indigo),
        useMaterial3: true,
      ),
      home: const Scaffold(
        body: Center(
          child: Text('Crypto Cloud Mining — Phase 0'),
        ),
      ),
    );
  }
}
