#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/pkt_cls.h>


SEC("classifier")
int handle_egress(struct __sk_buff *skb) {
    // Example logic to change destination MAC address
    // or perform other modifications

    // Mark the packet to be redirected to another interface (ifindex)
    int ifindex = 3;  // Example: Redirect to interface index 3
    bpf_redirect(ifindex, 0); // Redirects to another network interface

    return TC_ACT_REDIRECT;
}

char _license[] SEC("license") = "GPL";//
// Created by mjlong on 9/3/24.
//
