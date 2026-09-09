package metrics

// disk io creates the graph of
// number of reads and writes done

// the values is taken from
// {/proc/diskstats}
// all metrics are cumulative
// except field 9

// values to consider for the graph
// field 1- total number of reads completed.
// field 4 and 10 reads the miliseconds it spent reading and writing-- keep it for the latency graph
// field 5- total number of write completed.
// this metrics is the read and write operations done per sec

// field 3 and 7 gives the sectors written and read
// this is also cumulative
// these are the bytes read and written /sec
// the disk throughput is this

// identify how busy your disks are
