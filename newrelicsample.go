func StartDataSegmentNowForDataStore(op string, tableName string, txn newrelic.Transaction,datastore newrelic.DatastoreProduct) newrelic.DatastoreSegment {
		s := newrelic.DatastoreSegment{
			Product:    datastore,
			Collection: tableName,
			Operation:  op,
		}

		s.StartTime = newrelic.StartSegmentNow(txn)
		return s
	return newrelic.DatastoreSegment{}
}
