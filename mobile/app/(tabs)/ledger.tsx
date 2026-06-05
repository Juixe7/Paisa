import React from 'react';
import { View, Text, StyleSheet, FlatList, TouchableOpacity } from 'react-native';

const MOCK_TRANSACTIONS = [
  {
    id: 'tx_1',
    merchant: 'Swiggy',
    amount: 120.0,
    direction: 'debit',
    category: 'Dining Out',
    date: 'Today, 2:30 PM',
    source: 'SMS',
    confidence: 0.98,
  },
  {
    id: 'tx_2',
    merchant: 'RAMESH@paytm',
    amount: 350.0,
    direction: 'debit',
    category: 'Groceries',
    date: 'Yesterday, 11:15 AM',
    source: 'SMS',
    confidence: 0.75, // Low confidence -> Show yellow dot
  },
  {
    id: 'tx_3',
    merchant: 'Salary Credit',
    amount: 25000.0,
    direction: 'credit',
    category: 'Income',
    date: '1 Jun, 9:00 AM',
    source: 'Manual',
    confidence: 1.0,
  },
  {
    id: 'tx_4',
    merchant: 'Auto Fare',
    amount: 60.0,
    direction: 'debit',
    category: 'Commute',
    date: '30 May, 6:45 PM',
    source: 'Manual',
    confidence: 1.0,
  },
];

export default function LedgerScreen() {
  const renderItem = ({ item }: { item: typeof MOCK_TRANSACTIONS[0] }) => {
    const isDebit = item.direction === 'debit';
    const isLowConfidence = item.confidence < 0.8;

    return (
      <View style={styles.txRow}>
        <View style={styles.leftCol}>
          <View style={styles.merchantContainer}>
            {isLowConfidence && <View style={styles.yellowDot} />}
            <Text style={styles.merchantText}>{item.merchant}</Text>
          </View>
          <Text style={styles.detailsText}>
            {item.date} • {item.source}
          </Text>
        </View>

        <View style={styles.rightCol}>
          <Text style={[styles.amountText, isDebit ? styles.debitText : styles.creditText]}>
            {isDebit ? '-' : '+'}₹{item.amount.toFixed(2)}
          </Text>
          <TouchableOpacity style={styles.categoryBadge}>
            <Text style={styles.categoryText}>{item.category}</Text>
          </TouchableOpacity>
        </View>
      </View>
    );
  };

  return (
    <View style={styles.container}>
      <FlatList
        data={MOCK_TRANSACTIONS}
        keyExtractor={(item) => item.id}
        renderItem={renderItem}
        contentContainerStyle={styles.listContent}
        ItemSeparatorComponent={() => <View style={styles.separator} />}
        ListEmptyComponent={
          <View style={styles.emptyView}>
            <Text style={styles.emptyText}>No transactions recorded yet.</Text>
          </View>
        }
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#0F172A',
  },
  listContent: {
    padding: 16,
  },
  txRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingVertical: 12,
  },
  leftCol: {
    flex: 1,
  },
  merchantContainer: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  yellowDot: {
    width: 8,
    height: 8,
    borderRadius: 4,
    backgroundColor: '#F59E0B',
    marginRight: 8,
  },
  merchantText: {
    color: '#FFF',
    fontSize: 16,
    fontWeight: 'bold',
  },
  detailsText: {
    color: '#94A3B8',
    fontSize: 12,
    marginTop: 4,
  },
  rightCol: {
    alignItems: 'flex-end',
  },
  amountText: {
    fontSize: 16,
    fontWeight: 'bold',
  },
  debitText: {
    color: '#EF4444',
  },
  creditText: {
    color: '#10B981',
  },
  categoryBadge: {
    backgroundColor: '#1E293B',
    borderRadius: 6,
    paddingHorizontal: 8,
    paddingVertical: 4,
    marginTop: 6,
    borderWidth: 1,
    borderColor: '#334155',
  },
  categoryText: {
    color: '#38BDF8',
    fontSize: 11,
    fontWeight: '600',
  },
  separator: {
    height: 1,
    backgroundColor: '#1E293B',
    marginVertical: 4,
  },
  emptyView: {
    alignItems: 'center',
    padding: 40,
  },
  emptyText: {
    color: '#94A3B8',
  },
});
