//
// Created by mjlong on 9/3/24.
//
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/udp.h>
#include <linux/in.h>
#include <linux/pkt_cls.h>
//Copied from /usr/src/kernels/5.14.0-427.33.1.el9_4.x86_64/tools/lib/bpf/bpf_endian.hof which there were
//sevgeral versions but all that I checked were the same,
/**
/home/mjlong/go/src/github.com/mikejlong60/algorithms/skiena-4/bpf_endian.h
/home/mjlong/go/src/github.com/lizrice/learning-ebpf/libbpf/src/bpf_endian.h
/home/mjlong/go/src/github.com/lizrice/bpftool/libbpf/src/bpf_endian.h
/home/mjlong/go/src/github.com/lizrice/bpftool/src/libbpf/include/bpf/bpf_endian.h
/home/mjlong/go/src/github.com/lizrice/bpftool/src/bootstrap/libbpf/include/bpf/bpf_endian.h
/home/mjlong/go/src/github.com/libbpf/bpftool/libbpf/src/bpf_endian.h
/home/mjlong/go/src/github.com/libbpf/bpftool/src/libbpf/include/bpf/bpf_endian.h
/home/mjlong/go/src/github.com/libbpf/bpftool/src/bootstrap/libbpf/include/bpf/bpf_endian.h
find: ‘/run/user/1000/doc’: Permission denied
find: ‘/run/user/1000/gvfs’: Permission denied
/usr/include/bpf/bpf_endian.h
/usr/src/kernels/5.14.0-427.33.1.el9_4.x86_64/tools/bpf/resolve_btfids/libbpf/include/bpf/bpf_endian.h
/usr/src/kernels/5.14.0-427.33.1.el9_4.x86_64/tools/lib/bpf/bpf_endian.h
*/
#include "bpf_endian.h"


SEC("classifier")
int handle_egress(struct __sk_buff *skb) {
    void *data = (void *)(long)skb->data;
    void *data_end = (void *)(long)skb->data_end;

    // Parse Ethernet header
    struct ethhdr *eth = data;
    if (eth + 1 > data_end) {
        return TC_ACT_SHOT; // Drop packet if parsing fails
    }

    // Parse IP header
    struct iphdr *ip = data + sizeof(struct ethhdr);
    if (ip + 1 > data_end) {
        return TC_ACT_SHOT; // Drop packet if parsing fails
    }

    // Check if this is a UDP packet (for example)
    if (ip->protocol == IPPROTO_UDP) {
        // Modify IP destination address
        ip->daddr = bpf_htonl(0xC0A80102); // Example: New destination IP 192.168.1.2

        // Recalculate IP checksum (simplified)
        ip->check = 0;
        ip->check = bpf_csum_diff(0, 0, (__be32 *)ip, ip->ihl * 4, 0);

        // Redirect to another network interface (e.g., ifindex 3)
        return bpf_redirect(3, 0);
    }

    return TC_ACT_OK; // Allow packet to pass through unchanged
}

char _license[] SEC("license") = "GPL";