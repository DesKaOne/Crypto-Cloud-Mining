import 'package:flutter_test/flutter_test.dart';
import 'package:crypto_cloud_mining/main.dart';

void main() {
  testWidgets('phase 0 app boots', (tester) async {
    await tester.pumpWidget(const CryptoCloudMiningApp());
    expect(find.text('Crypto Cloud Mining — Phase 0'), findsOneWidget);
  });
}
