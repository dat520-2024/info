# Tentative Syllabus

The official textbook for the course is RSDP: Introduction to Reliable and Secure Distributed Programming, Christian Cachin, Rachid Guerraoui, and Luís Rodrigues, 2nd ed. Springer, 2011.
The following chapters are syllabus: 1, 2, 3, 4, 5, 6.

The sections that cover the log-recovery algorithms and Byzantine algorithms in Chapters 3-6 need not be considered for the exam.
All sections in 1 and 2 should be considered for the exam.

## The lecture presentations are syllabus

- [Course Introduction](slides2024/0-course-info-2024.pdf)
- [Introduction to Distributed Systems](slides2024/1-introduction-2024.pdf)
- [Gorums](slides2024/2-gorums-intro-2024.pdf)
- (will be added later) [Paxos I](slides2024/3-paxos-from-scratch-2024.pdf)
- (will be added later) [Paxos II](slides2024/4-paxos-made-insanely-simple-2024.pdf)
- [Designing for Understandability: The Raft Consensus Algorithm](https://www.youtube.com/watch?v=vYp4LYbnnW8) (video lecture)

The slides from the guest lectures are also syllabus; will be added above.

Optional video lectures related to the Raft paper:

- [Raft Lecture (Raft user study)](https://www.youtube.com/watch?v=YbZ3zDzDnrw)
- [Paxos Lecture (Paxos user study)](https://www.youtube.com/watch?v=JEpsBg0AO6o)

## The lab assignments available on GitHub are syllabus

## Paper Reading List

The following papers are required reading:

- [Paxos Explained from Scratch](reading/paxos-scratch-paper.pdf), Hein Meling and Leander Jehl
- [Paxos Made Simple](reading/paxos-simple.pdf), Leslie Lamport
- [In Search of an Understandable Consensus Algorithm](reading/raft.pdf), Diego Ongaro and John Ousterhout
- [Consistent Global States](reading/consistent-global-states.pdf), Keith Marzullo and Ozalp Babaoglu
- [Keeping CALM When Distributed Consistency Is Easy](reading/keeping-calm.pdf), Joseph M. Hellerstein, Peter Alvaro
- [The Tail at Scale](reading/tail-at-scale.pdf), Jeff Dean and Luiz André Barroso

Additional Paxos-related reading material (not required reading) can be found in lab4's [resources](../lab4/resources) folder.

The following wikipedia articles are syllabus:

- [The CAP Theorem](https://en.wikipedia.org/wiki/CAP_theorem)
- [The PACELC Theorem](https://en.wikipedia.org/wiki/PACELC_theorem)
- [The FLP Impossibility Result](<https://en.wikipedia.org/wiki/Consensus_(computer_science)#The_FLP_impossibility_result_for_asynchronous_deterministic_consensus>)
- [The Two Generals Problem](https://en.wikipedia.org/wiki/Two_Generals%27_Problem)
- [Fallacies of Distributed Computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing)
- [Byzantine Generals Problem](https://en.wikipedia.org/wiki/Byzantine_fault)
